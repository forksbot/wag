// Generated by internal/cmd/syscalls/generate.go

package main

import "encoding/binary"

func importRead() uint64
func importWrite() uint64
func importClose() uint64
func importLseek() uint64
func importPread() uint64
func importPwrite() uint64
func importDup() uint64
func importGetpid() uint64
func importSendfile() uint64
func importShutdown() uint64
func importSocketpair() uint64
func importFlock() uint64
func importFsync() uint64
func importFdatasync() uint64
func importTruncate() uint64
func importFtruncate() uint64
func importGetcwd() uint64
func importChdir() uint64
func importFchdir() uint64
func importFchmod() uint64
func importFchown() uint64
func importLchown() uint64
func importUmask() uint64
func importGetuid() uint64
func importGetgid() uint64
func importVhangup() uint64
func importSync() uint64
func importGettid() uint64
func importTime() uint64
func importPosixFadvise() uint64
func importExit() uint64
func importInotifyInit1() uint64
func importInotifyAddWatch() uint64
func importInotifyRmWatch() uint64
func importOpenat() uint64
func importMkdirat() uint64
func importFchownat() uint64
func importUnlinkat() uint64
func importRenameat() uint64
func importLinkat() uint64
func importSymlinkat() uint64
func importReadlinkat() uint64
func importFchmodat() uint64
func importFaccessat() uint64
func importSplice() uint64
func importTee() uint64
func importSyncFileRange() uint64
func importFallocate() uint64
func importEventfd() uint64
func importDup3() uint64
func importPipe2() uint64

func init() {
	importVector = make([]byte, 424)
	binary.LittleEndian.PutUint64(importVector[416:], 0x7fff0000)
	binary.LittleEndian.PutUint64(importVector[408:], importTrapHandler())
	binary.LittleEndian.PutUint64(importVector[400:], importRead())
	binary.LittleEndian.PutUint64(importVector[392:], importWrite())
	binary.LittleEndian.PutUint64(importVector[384:], importClose())
	binary.LittleEndian.PutUint64(importVector[376:], importLseek())
	binary.LittleEndian.PutUint64(importVector[368:], importPread())
	binary.LittleEndian.PutUint64(importVector[360:], importPwrite())
	binary.LittleEndian.PutUint64(importVector[352:], importDup())
	binary.LittleEndian.PutUint64(importVector[344:], importGetpid())
	binary.LittleEndian.PutUint64(importVector[336:], importSendfile())
	binary.LittleEndian.PutUint64(importVector[328:], importShutdown())
	binary.LittleEndian.PutUint64(importVector[320:], importSocketpair())
	binary.LittleEndian.PutUint64(importVector[312:], importFlock())
	binary.LittleEndian.PutUint64(importVector[304:], importFsync())
	binary.LittleEndian.PutUint64(importVector[296:], importFdatasync())
	binary.LittleEndian.PutUint64(importVector[288:], importTruncate())
	binary.LittleEndian.PutUint64(importVector[280:], importFtruncate())
	binary.LittleEndian.PutUint64(importVector[272:], importGetcwd())
	binary.LittleEndian.PutUint64(importVector[264:], importChdir())
	binary.LittleEndian.PutUint64(importVector[256:], importFchdir())
	binary.LittleEndian.PutUint64(importVector[248:], importFchmod())
	binary.LittleEndian.PutUint64(importVector[240:], importFchown())
	binary.LittleEndian.PutUint64(importVector[232:], importLchown())
	binary.LittleEndian.PutUint64(importVector[224:], importUmask())
	binary.LittleEndian.PutUint64(importVector[216:], importGetuid())
	binary.LittleEndian.PutUint64(importVector[208:], importGetgid())
	binary.LittleEndian.PutUint64(importVector[200:], importVhangup())
	binary.LittleEndian.PutUint64(importVector[192:], importSync())
	binary.LittleEndian.PutUint64(importVector[184:], importGettid())
	binary.LittleEndian.PutUint64(importVector[176:], importTime())
	binary.LittleEndian.PutUint64(importVector[168:], importPosixFadvise())
	binary.LittleEndian.PutUint64(importVector[160:], importExit())
	binary.LittleEndian.PutUint64(importVector[152:], importInotifyInit1())
	binary.LittleEndian.PutUint64(importVector[144:], importInotifyAddWatch())
	binary.LittleEndian.PutUint64(importVector[136:], importInotifyRmWatch())
	binary.LittleEndian.PutUint64(importVector[128:], importOpenat())
	binary.LittleEndian.PutUint64(importVector[120:], importMkdirat())
	binary.LittleEndian.PutUint64(importVector[112:], importFchownat())
	binary.LittleEndian.PutUint64(importVector[104:], importUnlinkat())
	binary.LittleEndian.PutUint64(importVector[96:], importRenameat())
	binary.LittleEndian.PutUint64(importVector[88:], importLinkat())
	binary.LittleEndian.PutUint64(importVector[80:], importSymlinkat())
	binary.LittleEndian.PutUint64(importVector[72:], importReadlinkat())
	binary.LittleEndian.PutUint64(importVector[64:], importFchmodat())
	binary.LittleEndian.PutUint64(importVector[56:], importFaccessat())
	binary.LittleEndian.PutUint64(importVector[48:], importSplice())
	binary.LittleEndian.PutUint64(importVector[40:], importTee())
	binary.LittleEndian.PutUint64(importVector[32:], importSyncFileRange())
	binary.LittleEndian.PutUint64(importVector[24:], importFallocate())
	binary.LittleEndian.PutUint64(importVector[16:], importEventfd())
	binary.LittleEndian.PutUint64(importVector[8:], importDup3())
	binary.LittleEndian.PutUint64(importVector[0:], importPipe2())
	importFuncs["read"] = importFunc{-3, 3}
	importFuncs["write"] = importFunc{-4, 3}
	importFuncs["close"] = importFunc{-5, 1}
	importFuncs["lseek"] = importFunc{-6, 3}
	importFuncs["pread"] = importFunc{-7, 4}
	importFuncs["pwrite"] = importFunc{-8, 4}
	importFuncs["dup"] = importFunc{-9, 1}
	importFuncs["getpid"] = importFunc{-10, 0}
	importFuncs["sendfile"] = importFunc{-11, 4}
	importFuncs["shutdown"] = importFunc{-12, 2}
	importFuncs["socketpair"] = importFunc{-13, 4}
	importFuncs["flock"] = importFunc{-14, 2}
	importFuncs["fsync"] = importFunc{-15, 1}
	importFuncs["fdatasync"] = importFunc{-16, 1}
	importFuncs["truncate"] = importFunc{-17, 2}
	importFuncs["ftruncate"] = importFunc{-18, 2}
	importFuncs["getcwd"] = importFunc{-19, 2}
	importFuncs["chdir"] = importFunc{-20, 1}
	importFuncs["fchdir"] = importFunc{-21, 1}
	importFuncs["fchmod"] = importFunc{-22, 2}
	importFuncs["fchown"] = importFunc{-23, 3}
	importFuncs["lchown"] = importFunc{-24, 3}
	importFuncs["umask"] = importFunc{-25, 1}
	importFuncs["getuid"] = importFunc{-26, 0}
	importFuncs["getgid"] = importFunc{-27, 0}
	importFuncs["vhangup"] = importFunc{-28, 0}
	importFuncs["sync"] = importFunc{-29, 0}
	importFuncs["gettid"] = importFunc{-30, 0}
	importFuncs["time"] = importFunc{-31, 1}
	importFuncs["posix_fadvise"] = importFunc{-32, 4}
	importFuncs["_exit"] = importFunc{-33, 1}
	importFuncs["inotify_init1"] = importFunc{-34, 0}
	importFuncs["inotify_add_watch"] = importFunc{-35, 3}
	importFuncs["inotify_rm_watch"] = importFunc{-36, 2}
	importFuncs["openat"] = importFunc{-37, 4}
	importFuncs["mkdirat"] = importFunc{-38, 3}
	importFuncs["fchownat"] = importFunc{-39, 5}
	importFuncs["unlinkat"] = importFunc{-40, 3}
	importFuncs["renameat"] = importFunc{-41, 4}
	importFuncs["linkat"] = importFunc{-42, 5}
	importFuncs["symlinkat"] = importFunc{-43, 3}
	importFuncs["readlinkat"] = importFunc{-44, 4}
	importFuncs["fchmodat"] = importFunc{-45, 4}
	importFuncs["faccessat"] = importFunc{-46, 4}
	importFuncs["splice"] = importFunc{-47, 6}
	importFuncs["tee"] = importFunc{-48, 4}
	importFuncs["sync_file_range"] = importFunc{-49, 4}
	importFuncs["fallocate"] = importFunc{-50, 4}
	importFuncs["eventfd"] = importFunc{-51, 2}
	importFuncs["dup3"] = importFunc{-52, 3}
	importFuncs["pipe2"] = importFunc{-53, 2}
}
