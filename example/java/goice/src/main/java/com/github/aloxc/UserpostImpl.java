package com.github.aloxc;

import Ice.Current;
import com.github.aloxc.user.post._UserpostDisp;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Map;
import java.util.Random;

public class UserpostImpl extends _UserpostDisp {
    private static final Logger logger = LoggerFactory.getLogger(UserpostImpl.class);

    @Override
    public String valuestr(String name, String value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},name = {},value = {}", method, ctx, name,value);
        return "name = " + name + ",value = " + value;
    }

    @Override
    public String valuelong(String name, long value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},name = {},value = {}", method, ctx, name,value);
        return "name = " + name + ",long = " + value;
    }

    @Override
    public String threeparams(String name, int value, double any, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},name = {},value = {},any = {}", method, ctx, name,value,any);
        return null;
    }

    @Override
    public String doarr(String[] sarr, int[] iarr, String key, String value, int i, Current current) {
        StringBuffer buffer = new StringBuffer();
        if (sarr != null){
            for (int j = 0; j < sarr.length; j++) {
                buffer.append("sarr[");
                buffer.append(j);
                buffer.append("] = " );
                buffer.append(sarr[j]);
                buffer.append("\n");
            }
        }
        buffer.append("\n");
        if (iarr != null){
            for (int j = 0; j < iarr.length; j++) {
                buffer.append("iarr[");
                buffer.append(j);
                buffer.append("] = " );
                buffer.append(iarr[j]);
                buffer.append("\n");
            }
        }
        buffer.append("\n");
        buffer.append("i=");
        buffer.append(i);

        buffer.append("\n");
        buffer.append("key=");
        buffer.append(key);

        buffer.append("\n");
        buffer.append("value=");
        buffer.append(value);
        return buffer.toString();
    }

    @Override
    public String todo(String json, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},json = {}", method, ctx, json);
        return json;
    }

    @Override
    public int[] getIntArr(int i, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},i = {}", method, ctx, i);
        int[] arr = new int[i];
        for (int j = 0; j < i; j++) {
            arr[j] = j + 500;
        }
        return arr;
    }

    @Override
    public String[] getStrArr(int i, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        String[] arr = new String[i];
        Random r = new Random();
        for (int j = 0; j < i; j++) {
            arr[j] = "a" + (j + 500);
        }
        StringBuffer buffer = new StringBuffer();
        for (int j = 0; j < 10; j++) {
            buffer.append("aaabbbcccdddeeefffggghhhiiijjj");//添加了三十个
        }
        int x= r.nextInt(i);
        logger.info("方法[{}],ctx={},i = {},x = {}", method, ctx, i,x);
        arr[x] = buffer.toString();
        return arr;
    }

    public int[] getIntArr(Current current) {
        return new int[]{5,6,7,9};
    }

    public String[] getStrArr(Current current) {
        return new String[]{"ax","by","cz"};
    }
}