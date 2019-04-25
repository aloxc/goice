[["java:package:com.github.aloxc"]]  
module user{
	module post{
		sequence<string> StrArr;
		sequence<int> IntArr;

		interface Userpost{
			string valuestr(string name,string value);
			string valuelong(string name,long value);
			string threeparams(string name,int value,double any);
			string doarr(StrArr sar,IntArr iarr,string key,string value,int i);
			string todo(string json);
			IntArr getIntArr(int i);
			StrArr getStrArr(int i);
		};
	};
};
