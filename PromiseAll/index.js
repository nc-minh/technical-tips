const X = (x) => {
	return new Promise((resolve) => {
		setTimeout(() => {
			resolve(x);
		}, 1000);
	});
};

const Y = (y) => {
	return new Promise((resolve) => {
		setTimeout(() => {
			resolve(y);
		}, 3000);
	});
};

const total = async (x, y) => {
	console.time("Timer");

	// const _x = await X(x);
	// const _y = await Y(y);

	const [_x, _y] = await Promise.all([X(x), Y(y)]);

	console.timeEnd("Timer");

	return _x + _y;
};

total(1, 2).then((result) => console.log(result));
