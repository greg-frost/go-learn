module main

go 1.19

replace (
	ypmodule => ../ypmodule
	workspace/user/repo => ../somemodule
)

require (
	workspace/user/repo v0.0.0-00010101000000-000000000000 // indirect
	ypmodule v0.0.0-00010101000000-000000000000 // indirect
)
