package ice

type Incoming struct {

}
//响应头部10字节 49 63 65 50 01 00 01 00 02 00  // 10
//数据总长度， 1D 00 00 00   //10+4 = 14
//请求id   02 00 00 00   //14+4 = 18
// 一个byte的零 00       //18+1=19   java代码ReplyStatus.replyOK,如果异常就是ReplyStatus.replyUserException
// 一个int的零 数据长度，最终会整形的（整形是总的数据长度去掉这个位置之前的长度） 0A 00 00 00         //19 + 4 = 23
// 加上encodingVersion  01 01     //23+2 = 25
//实际数据
//数据长度  03
//数据   62 62 62