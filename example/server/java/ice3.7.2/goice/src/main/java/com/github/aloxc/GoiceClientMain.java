package com.github.aloxc;

import com.zeroc.Ice.*;
import com.github.aloxc.goiceinter.GoicePrx;
import org.slf4j.LoggerFactory;

import java.util.Random;
import java.util.concurrent.atomic.AtomicInteger;

public class GoiceClientMain {

    public static AtomicInteger COUNT = new AtomicInteger();
    private static org.slf4j.Logger logger = LoggerFactory.getLogger(GoiceClientMain.class);

    public static void main(String[] args) throws InterruptedException, NoSuchFieldException, IllegalAccessException {
        // 通信器
        Communicator ic;
        ic = Util.initialize(args);
        ObjectPrx proxy = ic.stringToProxy("Goice:default -p 1888");
        GoicePrx goicePrx = GoicePrx.checkedCast(proxy);
//        Thread.sleep(10000);
        Random random = new Random();
//        HashMap<String,String> context = new HashMap<>();
//        context.put("a","aloxc");
        long start = System.currentTimeMillis();
        int times = 100000;
        for (int i = 1; i < times; i++) {

            goicePrx.two("我","你");
//            if(i%2000==0){
//                System.out.println(i+"," + (System.currentTimeMillis() - start));
//            }
        }
        System.out.printf("执行[%d]花费[%d]",times,System.currentTimeMillis() - start);

//        System.out.println(goicePrx.getAge(10));
//        System.out.println(goicePrx.getArticle("free", 124));
//        Request request = new Request();
//        request.method = "exception";
//        HashMap<String, String> params = new HashMap<>();
//        params.put("item", "free");
//        params.put("id", "111");
//        request.params = params;
//        System.out.println(goicePrx.execute(request));

        ic.destroy();


    }
}
