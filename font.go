package captcha

const (
	numberWidth  = 5
	numberHeight = 8
)

var font = [][]byte{
	{ // 0
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 1
		0, 0, 1, 0, 0,
		0, 1, 1, 0, 0,
		1, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		1, 1, 1, 1, 1,
	},
	{ // 2
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 1,
		0, 1, 1, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 1,
	},
	{ // 3
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
	{ // 4
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 1, 1, 1, 1,
		0, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
	},
	{ // 5
		1, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
	{ // 6
		0, 0, 1, 1, 1,
		0, 1, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 7
		1, 1, 1, 1, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
	},
	{ // 8
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 9
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 1, 0, 0, 1,
		0, 1, 1, 1, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
}
