const obj = {
	name: "Minh",
	address: {
		name: "HaNoi",
	},
	age: 18,
	salary: undefined,
};

// const copyObj = obj; //shallow copy

// const copyObj = {...obj}; //shallow copy

// const copyObj = JSON.parse(JSON.stringify(obj)); //shallow copy

const copyObj = structuredClone(obj);// nodejs >= 17.x.x

obj.name = "Cong Minh";
obj.address.name = "Ha Noi - Viet Nam";

console.log(`Obj =>> ${obj}`);
console.log(`CopyObj =>> ${copyObj}`);
