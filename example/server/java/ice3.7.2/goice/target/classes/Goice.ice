[["java:package:com.github.aloxc"]]  
module goiceinter{ 
	const string Name = "\u0067\u006f\u0069\u0063\u0065\u793a\u4f8b\u0069\u0063\u0065\u63a5\u53e3\u6587\u4ef6";
	enum Color{Red,Blue,Green,Yello};
	enum HouseType{Apartment = 5,Villa};
	dictionary<string, string> Params;
	sequence<string> StrArr;
	sequence<int> IntArr;
	sequence<bool> BoolArr;
	sequence<byte> ByteArr;
	sequence<short> ShortArr;
	sequence<long> LongArr;
	sequence<float> FloatArr;
	sequence<double> DoubleArr;
	struct Request{
		string method;
		Params params;
	};



    interface Goice{
	string non();
        string getString();
        string getStringFrom(string value);
	StrArr getStringArr();
        StrArr getStringArrFrom(StrArr arr);
	string two(string from,string to);
	void vvoid();
	void vvoidTo(string value);
	byte getByte();
	byte getByteFrom(byte value);
	ByteArr getByteArr();
	ByteArr getByteArrFrom(ByteArr arr);
	bool getBool();
	bool getBoolFrom(bool value);
	BoolArr getBoolArr();
	BoolArr getBoolArrFrom(BoolArr arr);
	short getShort();
	short getShortFrom(short value);
	ShortArr getShortArr();
	ShortArr getShortArrFrom(ShortArr arr);
	int getInt();
	int getIntFrom(int value);
	IntArr getIntArr();
	IntArr getIntArrFrom(IntArr arr);
	long getLong();
	long getLongFrom(long value);
	LongArr getLongArr();
	LongArr getLongArrFrom(LongArr arr);
	float getFloat();
	float getFloatFrom(float value);
	FloatArr getFloatArr();
	FloatArr getFloatArrFrom(FloatArr arr);
	double getDouble();
	double getDoubleFrom(double value);
	DoubleArr getDoubleArr();
	DoubleArr getDoubleArrFrom(DoubleArr arr);
	string execute(Request request);

    };

};