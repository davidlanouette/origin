// +build linux

// Tests for public API of libseccomp Go bindings

package seccomp

import (
	"fmt"
	"syscall"
	"testing"
)

// Type Function Tests

type versionErrorTest struct {
	err VersionError
	str string
}

var versionStr = fmt.Sprintf("%d.%d.%d", verMajor, verMinor, verMicro)

var versionErrorTests = []versionErrorTest{
	{
		VersionError{
			"deadbeef",
			"x.y.z",
		},
		"Libseccomp version too low: deadbeef: " +
			"minimum supported is x.y.z: detected " + versionStr,
	},
	{
		VersionError{
			"",
			"x.y.z",
		},
		"Libseccomp version too low: minimum supported is x.y.z: " +
			"detected " + versionStr,
	},
	{
		VersionError{
			"deadbeef",
			"",
		},
		"Libseccomp version too low: " +
			"deadbeef: minimum supported is 2.2.0: " +
			"detected " + versionStr,
	},
	{
		VersionError{
			"",
			"",
		},
		"Libseccomp version too low: minimum supported is 2.2.0: " +
			"detected " + versionStr,
	},
}

func TestVersionError(t *testing.T) {
	for i, test := range versionErrorTests {
		str := test.err.Error()
		if str != test.str {
			t.Errorf("VersionError %d: got %q: expected %q", i, str, test.str)
		}
	}
}

func ApiLevelIsSupported() bool {
	return verMajor > 2 ||
		(verMajor == 2 && verMinor > 3) ||
		(verMajor == 2 && verMinor == 3 && verMicro >= 3)
}

func TestGetApiLevel(t *testing.T) {
	api, err := GetApi()
	if !ApiLevelIsSupported() {
		if api != 0 {
			t.Errorf("API level returned despite lack of support: %v", api)
		} else if err == nil {
			t.Errorf("No error returned despite lack of API level support")
		}

		t.Skipf("Skipping test: %s", err)
	} else if err != nil {
		t.Errorf("Error getting API level: %s", err)
	}
	fmt.Printf("Got API level of %v\n", api)
}

func TestSetApiLevel(t *testing.T) {
	var expectedApi uint

	expectedApi = 1
	err := SetApi(expectedApi)
	if !ApiLevelIsSupported() {
		if err == nil {
			t.Errorf("No error returned despite lack of API level support")
		}

		t.Skipf("Skipping test: %s", err)
	} else if err != nil {
		t.Errorf("Error setting API level: %s", err)
	}

	api, err := GetApi()
	if err != nil {
		t.Errorf("Error getting API level: %s", err)
	} else if api != expectedApi {
		t.Errorf("Got API level %v: expected %v", api, expectedApi)
	}
}

func TestActionSetReturnCode(t *testing.T) {
	if ActInvalid.SetReturnCode(0x0010) != ActInvalid {
		t.Errorf("Able to set a return code on invalid action!")
	}

	codeSet := ActErrno.SetReturnCode(0x0001)
	if codeSet == ActErrno || codeSet.GetReturnCode() != 0x0001 {
		t.Errorf("Could not set return code on ActErrno")
	}
}

func TestSyscallGetName(t *testing.T) {
	call1 := ScmpSyscall(0x1)
	callFail := ScmpSyscall(0x999)

	name, err := call1.GetName()
	if err != nil {
		t.Errorf("Error getting syscall name for number 0x1")
	} else if len(name) == 0 {
		t.Errorf("Empty name returned for syscall 0x1")
	}
	fmt.Printf("Got name of syscall 0x1 on native arch as %s\n", name)

	_, err = callFail.GetName()
	if err == nil {
		t.Errorf("Getting nonexistant syscall should error!")
	}
}

func TestSyscallGetNameByArch(t *testing.T) {
	call1 := ScmpSyscall(0x1)
	callInvalid := ScmpSyscall(0x999)
	archGood := ArchAMD64
	archBad := ArchInvalid

	name, err := call1.GetNameByArch(archGood)
	if err != nil {
		t.Errorf("Error getting syscall name for number 0x1 and arch AMD64")
	} else if name != "write" {
		t.Errorf("Got incorrect name for syscall 0x1 - expected write, got %s", name)
	}

	_, err = call1.GetNameByArch(archBad)
	if err == nil {
		t.Errorf("Bad architecture GetNameByArch() should error!")
	}

	_, err = callInvalid.GetNameByArch(archGood)
	if err == nil {
		t.Errorf("Bad syscall GetNameByArch() should error!")
	}

	_, err = callInvalid.GetNameByArch(archBad)
	if err == nil {
		t.Errorf("Bad syscall and bad arch GetNameByArch() should error!")
	}
}

func TestGetSyscallFromName(t *testing.T) {
	name1 := "write"
	nameInval := "NOTASYSCALL"

	syscall, err := GetSyscallFromName(name1)
	if err != nil {
		t.Errorf("Error getting syscall number of write: %s", err)
	}
	fmt.Printf("Got syscall number of write on native arch as %d\n", syscall)

	_, err = GetSyscallFromName(nameInval)
	if err == nil {
		t.Errorf("Getting an invalid syscall should error!")
	}
}

func TestGetSyscallFromNameByArch(t *testing.T) {
	name1 := "write"
	nameInval := "NOTASYSCALL"
	arch1 := ArchAMD64
	archInval := ArchInvalid

	syscall, err := GetSyscallFromNameByArch(name1, arch1)
	if err != nil {
		t.Errorf("Error getting syscall number of write on AMD64: %s", err)
	}
	fmt.Printf("Got syscall number of write on AMD64 as %d\n", syscall)

	_, err = GetSyscallFromNameByArch(nameInval, arch1)
	if err == nil {
		t.Errorf("Getting invalid syscall with valid arch should error")
	}

	_, err = GetSyscallFromNameByArch(name1, archInval)
	if err == nil {
		t.Errorf("Getting valid syscall for invalid arch should error")
	}

	_, err = GetSyscallFromNameByArch(nameInval, archInval)
	if err == nil {
		t.Errorf("Getting invalid syscall for invalid arch should error")
	}
}

func TestMakeCondition(t *testing.T) {
	condition, err := MakeCondition(3, CompareNotEqual, 0x10)
	if err != nil {
		t.Errorf("Error making condition struct: %s", err)
	} else if condition.Argument != 3 || condition.Operand1 != 0x10 ||
		condition.Operand2 != 0 || condition.Op != CompareNotEqual {
		t.Errorf("Condition struct was filled incorrectly")
	}

	condition, err = MakeCondition(3, CompareMaskedEqual, 0x10, 0x20)
	if err != nil {
		t.Errorf("Error making condition struct: %s", err)
	} else if condition.Argument != 3 || condition.Operand1 != 0x10 ||
		condition.Operand2 != 0x20 || condition.Op != CompareMaskedEqual {
		t.Errorf("Condition struct was filled incorrectly")
	}

	_, err = MakeCondition(7, CompareNotEqual, 0x10)
	if err == nil {
		t.Errorf("Condition struct with bad syscall argument number should error")
	}

	_, err = MakeCondition(3, CompareInvalid, 0x10)
	if err == nil {
		t.Errorf("Condition struct with bad comparison operator should error")
	}

	_, err = MakeCondition(3, CompareMaskedEqual, 0x10, 0x20, 0x30)
	if err == nil {
		t.Errorf("MakeCondition with more than 2 arguments should fail")
	}

	_, err = MakeCondition(3, CompareMaskedEqual)
	if err == nil {
		t.Errorf("MakeCondition with no arguments should fail")
	}
}

// Utility Function Tests

func TestGetNativeArch(t *testing.T) {
	arch, err := GetNativeArch()
	if err != nil {
		t.Errorf("GetNativeArch should not error!")
	}
	fmt.Printf("Got native arch of system as %s\n", arch.String())
}

// Filter Tests

func TestFilterCreateRelease(t *testing.T) {
	_, err := NewFilter(ActInvalid)
	if err == nil {
		t.Errorf("Can create filter with invalid action")
	}

	filter, err := NewFilter(ActKill)
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}

	if !filter.IsValid() {
		t.Errorf("Filter created by NewFilter was not valid")
	}

	filter.Release()

	if filter.IsValid() {
		t.Errorf("Filter is valid after being released")
	}
}

func TestFilterReset(t *testing.T) {
	filter, err := NewFilter(ActKill)
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}
	defer filter.Release()

	// Ensure the default action is ActKill
	action, err := filter.GetDefaultAction()
	if err != nil {
		t.Errorf("Error getting default action of filter")
	} else if action != ActKill {
		t.Errorf("Default action of filter was set incorrectly!")
	}

	// Reset with a different default action
	err = filter.Reset(ActAllow)
	if err != nil {
		t.Errorf("Error resetting filter!")
	}

	valid := filter.IsValid()
	if !valid {
		t.Errorf("Filter is no longer valid after reset!")
	}

	// The default action should no longer be ActKill
	action, err = filter.GetDefaultAction()
	if err != nil {
		t.Errorf("Error getting default action of filter")
	} else if action != ActAllow {
		t.Errorf("Default action of filter was set incorrectly!")
	}
}

func TestFilterArchFunctions(t *testing.T) {
	filter, err := NewFilter(ActKill)
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}
	defer filter.Release()

	arch, err := GetNativeArch()
	if err != nil {
		t.Errorf("Error getting native architecture: %s", err)
	}

	present, err := filter.IsArchPresent(arch)
	if err != nil {
		t.Errorf("Error retrieving arch from filter: %s", err)
	} else if !present {
		t.Errorf("Filter does not contain native architecture by default")
	}

	// Adding the native arch again should succeed, as it's already present
	err = filter.AddArch(arch)
	if err != nil {
		t.Errorf("Adding arch to filter already containing it should succeed")
	}

	// Make sure we don't add the native arch again
	prospectiveArch := ArchX86
	if arch == ArchX86 {
		prospectiveArch = ArchAMD64
	}

	// Check to make sure this other arch isn't in the filter
	present, err = filter.IsArchPresent(prospectiveArch)
	if err != nil {
		t.Errorf("Error retrieving arch from filter: %s", err)
	} else if present {
		t.Errorf("Arch not added to filter is present")
	}

	// Try removing the nonexistant arch - should succeed
	err = filter.RemoveArch(prospectiveArch)
	if err != nil {
		t.Errorf("Error removing nonexistant arch: %s", err)
	}

	// Add an arch, see if it's in the filter
	err = filter.AddArch(prospectiveArch)
	if err != nil {
		t.Errorf("Could not add arch %s to filter: %s",
			prospectiveArch.String(), err)
	}

	present, err = filter.IsArchPresent(prospectiveArch)
	if err != nil {
		t.Errorf("Error retrieving arch from filter: %s", err)
	} else if !present {
		t.Errorf("Filter does not contain architecture %s after it was added",
			prospectiveArch.String())
	}

	// Remove the arch again, make sure it's not in the filter
	err = filter.RemoveArch(prospectiveArch)
	if err != nil {
		t.Errorf("Could not remove arch %s from filter: %s",
			prospectiveArch.String(), err)
	}

	present, err = filter.IsArchPresent(prospectiveArch)
	if err != nil {
		t.Errorf("Error retrieving arch from filter: %s", err)
	} else if present {
		t.Errorf("Filter contains architecture %s after it was removed",
			prospectiveArch.String())
	}
}

func TestFilterAttributeGettersAndSetters(t *testing.T) {
	filter, err := NewFilter(ActKill)
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}
	defer filter.Release()

	act, err := filter.GetDefaultAction()
	if err != nil {
		t.Errorf("Error getting default action: %s", err)
	} else if act != ActKill {
		t.Errorf("Default action was set incorrectly")
	}

	err = filter.SetBadArchAction(ActAllow)
	if err != nil {
		t.Errorf("Error setting bad arch action: %s", err)
	}

	act, err = filter.GetBadArchAction()
	if err != nil {
		t.Errorf("Error getting bad arch action")
	} else if act != ActAllow {
		t.Errorf("Bad arch action was not set correcly!")
	}

	err = filter.SetNoNewPrivsBit(false)
	if err != nil {
		t.Errorf("Error setting no new privileges bit")
	}

	privs, err := filter.GetNoNewPrivsBit()
	if err != nil {
		t.Errorf("Error getting no new privileges bit!")
	} else if privs != false {
		t.Errorf("No new privileges bit was not set correctly")
	}

	if ApiLevelIsSupported() {
		api, err := GetApi()
		if err != nil {
			t.Errorf("Error getting API level: %s", err)
		} else if api < 3 {
			err = SetApi(3)
			if err != nil {
				t.Errorf("Error setting API level: %s", err)
			}
		}
	}

	err = filter.SetLogBit(true)
	if err != nil {
		if !ApiLevelIsSupported() {
			t.Logf("Ignoring failure: %s\n", err)
		} else {
			t.Errorf("Error setting log bit")
		}
	}

	log, err := filter.GetLogBit()
	if err != nil {
		if !ApiLevelIsSupported() {
			t.Logf("Ignoring failure: %s\n", err)
		} else {
			t.Errorf("Error getting log bit")
		}
	} else if log != true {
		t.Errorf("Log bit was not set correctly")
	}

	err = filter.SetBadArchAction(ActInvalid)
	if err == nil {
		t.Errorf("Setting bad arch action to an invalid action should error")
	}
}

func TestMergeFilters(t *testing.T) {
	filter1, err := NewFilter(ActAllow)
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}

	filter2, err := NewFilter(ActAllow)
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}

	// Need to remove the native arch and add another to the second filter
	// Filters must NOT share architectures to be successfully merged
	nativeArch, err := GetNativeArch()
	if err != nil {
		t.Errorf("Error getting native arch: %s", err)
	}

	prospectiveArch := ArchAMD64
	if nativeArch == ArchAMD64 {
		prospectiveArch = ArchX86
	}

	err = filter2.AddArch(prospectiveArch)
	if err != nil {
		t.Errorf("Error adding architecture to filter: %s", err)
	}

	err = filter2.RemoveArch(nativeArch)
	if err != nil {
		t.Errorf("Error removing architecture from filter: %s", err)
	}

	err = filter1.Merge(filter2)
	if err != nil {
		t.Errorf("Error merging filters: %s", err)
	}

	if filter2.IsValid() {
		t.Errorf("Source filter should not be valid after merging")
	}

	filter3, err := NewFilter(ActKill)
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}
	defer filter3.Release()

	err = filter1.Merge(filter3)
	if err == nil {
		t.Errorf("Attributes should have to match to merge filters")
	}
}

func TestRuleAddAndLoad(t *testing.T) {
	// Test #1: Add a trivial filter
	filter1, err := NewFilter(ActAllow)
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}
	defer filter1.Release()

	call, err := GetSyscallFromName("getpid")
	if err != nil {
		t.Errorf("Error getting syscall number of getpid: %s", err)
	}

	call2, err := GetSyscallFromName("setreuid")
	if err != nil {
		t.Errorf("Error getting syscall number of setreuid: %s", err)
	}

	call3, err := GetSyscallFromName("setreuid32")
	if err != nil {
		t.Errorf("Error getting syscall number of setreuid32: %s", err)
	}

	uid := syscall.Getuid()
	euid := syscall.Geteuid()

	err = filter1.AddRule(call, ActErrno.SetReturnCode(0x1))
	if err != nil {
		t.Errorf("Error adding rule to restrict syscall: %s", err)
	}

	cond, err := MakeCondition(1, CompareEqual, uint64(euid))
	if err != nil {
		t.Errorf("Error making rule to restrict syscall: %s", err)
	}

	cond2, err := MakeCondition(0, CompareEqual, uint64(uid))
	if err != nil {
		t.Errorf("Error making rule to restrict syscall: %s", err)
	}

	conditions := []ScmpCondition{cond, cond2}

	err = filter1.AddRuleConditional(call2, ActErrno.SetReturnCode(0x2), conditions)
	if err != nil {
		t.Errorf("Error adding conditional rule: %s", err)
	}

	err = filter1.AddRuleConditional(call3, ActErrno.SetReturnCode(0x3), conditions)
	if err != nil {
		t.Errorf("Error adding second conditional rule: %s", err)
	}

	err = filter1.Load()
	if err != nil {
		t.Errorf("Error loading filter: %s", err)
	}

	// Try making a simple syscall, it should error
	pid := syscall.Getpid()
	if pid != -1 {
		t.Errorf("Syscall should have returned error code!")
	}

	// Try making a Geteuid syscall that should normally succeed
	err = syscall.Setreuid(uid, euid)
	if err == nil {
		t.Errorf("Syscall should have returned error code!")
	} else if err != syscall.Errno(2) && err != syscall.Errno(3) {
		t.Errorf("Syscall returned incorrect error code - likely not blocked by Seccomp!")
	}
}

func TestLogAct(t *testing.T) {
	expectedPid := syscall.Getpid()

	api, err := GetApi()
	if err != nil {
		if !ApiLevelIsSupported() {
			t.Skipf("Skipping test: %s", err)
		}

		t.Errorf("Error getting API level: %s", err)
	} else if api < 3 {
		t.Skipf("Skipping test: API level %d is less than 3", api)
	}

	filter, err := NewFilter(ActErrno.SetReturnCode(0x0001))
	if err != nil {
		t.Errorf("Error creating filter: %s", err)
	}
	defer filter.Release()

	call, err := GetSyscallFromName("getpid")
	if err != nil {
		t.Errorf("Error getting syscall number of getpid: %s", err)
	}

	call1, err := GetSyscallFromName("write")
	if err != nil {
		t.Errorf("Error getting syscall number of write: %s", err)
	}

	call2, err := GetSyscallFromName("futex")
	if err != nil {
		t.Errorf("Error getting syscall number of futex: %s", err)
	}

	call3, err := GetSyscallFromName("exit_group")
	if err != nil {
		t.Errorf("Error getting syscall number of exit_group: %s", err)
	}

	err = filter.AddRule(call, ActLog)
	if err != nil {
		t.Errorf("Error adding rule to log syscall: %s", err)
	}

	err = filter.AddRule(call1, ActAllow)
	if err != nil {
		t.Errorf("Error adding rule to allow write syscall: %s", err)
	}

	err = filter.AddRule(call2, ActAllow)
	if err != nil {
		t.Errorf("Error adding rule to allow futex syscall: %s", err)
	}

	err = filter.AddRule(call3, ActAllow)
	if err != nil {
		t.Errorf("Error adding rule to allow exit_group syscall: %s", err)
	}

	err = filter.Load()
	if err != nil {
		t.Errorf("Error loading filter: %s", err)
	}

	// Try making a simple syscall, it should succeed
	pid := syscall.Getpid()
	if pid != expectedPid {
		t.Errorf("Syscall should have returned expected pid (%d != %d)", pid, expectedPid)
	}
}
