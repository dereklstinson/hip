# hipgo
go bindings for hip

This package might not work out of the box.  Infact, I had to fork hip in order to get it to work.
As of today 10/29/2019. This should work right out of the box (after installing rocm).



## Pre-use stuff

to install hip on pc go to https://rocm.github.io/install.html.

use 
```
go get github.com/dereklstinson/cutil
go get github.com/dereklstinson/half
```



## Examples
Look in the kernel folder for example on how to build a kernel.

Look in the testhip folder to see how it is done in c.

Look in the testmain to see how to use package.  

## Tips

I haven't tested this using goroutines. Other than you might want to use runtime.LockHostThread() when inside a goroutine.

GPUs are meant to run concurrently with the cpu.  You will need to run stream.sync() if something needs the result of a gpu function.

## Notes

This doesn't support modules for now, because it sucks for development using the Go extention tools on vscode.

This is not the whole hip_runtime api.  If you look through some of the files. I have commented functions out that
need to be binded. (yes, I bind by hand)

I might not work on this for a while. So, if anything is needed please send me a pull request.








