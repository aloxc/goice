package com.github.aloxc;

import Ice.Current;
import com.github.aloxc.goiceinter.Request;
import com.github.aloxc.goiceinter._GoiceDisp;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Random;
import java.util.concurrent.TimeUnit;

public class GoiceImpl extends _GoiceDisp {
    private static final Logger logger = LoggerFactory.getLogger(GoiceImpl.class);


    @Override
    public String non(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return method;
    }

    @Override
    public String getString(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return null;
    }

    @Override
    public String getStringFrom(String value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx, value);
        return method;
    }

    @Override
    public String[] getStringArr(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        Random r = new Random();
        int i = 100;
        String[] arr = new String[i];
        for (int j = 0; j < i; j++) {
            arr[j] = "a" + (j + 500);
        }
        StringBuffer buffer = new StringBuffer();
        for (int j = 0; j < 10; j++) {
            buffer.append("aaabbbcccdddeeefffggghhhiiijjj");//添加了三十个
        }
        int x = r.nextInt(i);
        logger.info("方法[{}],ctx={},i = {},x = {}", method, ctx, i, x);
        arr[x] = buffer.toString();
        return arr;
    }

    @Override
    public String[] getStringArrFrom(String[] arr, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},arr = {}", method, ctx, Arrays.toString(arr));
        return arr;
    }

    @Override
    public String two(String from, String to, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},from = {},to = {}", method, ctx, from, to);
        return "from = " + from + " to = " + to;
    }

    @Override
    public void vvoid(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
    }

    @Override
    public void vvoidTo(String to, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={} , to = {}", method, ctx, to);
    }

    @Override
    public byte getByte(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        Random r = new Random();
        int i = r.nextInt(100);
        return i % 2 == 0 ? (byte)(0 - i) : (byte)i;
    }

    @Override
    public byte getByteFrom(byte value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx, value);
        return (byte)(value + 1);
    }

    @Override
    public byte[] getByteArr(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={} ", method, ctx);
        byte[] bytes = new byte[]{'a','b'};
        return bytes;
    }

    @Override
    public byte[] getByteArrFrom(byte[] arr, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return arr;
    }


    @Override
    public boolean getBool(Current current) {

        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        Random r = new Random();
        return r.nextInt() % 2 == 0;    }

    @Override
    public boolean getBoolFrom(boolean bo, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},bo = {}", method, ctx, bo);
        Random r = new Random();
        return r.nextInt() % 2 == 0;
    }

    @Override
    public boolean[] getBoolArr(Current current) {
        Random r = new Random();
        int size = 10;
        boolean[] arr = new boolean[size];
        for (int i = 0; i < size; i++) {
            arr[i] = i %2 ==0;
        }
        arr[size - 1] = true;
        return arr;
    }

    @Override
    public boolean[] getBoolArrFrom(boolean[] arr, Current current) {
        for (int i = 0; i < arr.length; i++) {
            arr[i] = !arr[i];
        }
        return arr;
    }


    @Override
    public short getShort(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return -32;
    }

    @Override
    public short getShortFrom(short value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx, value);
        return value += 5;
    }

    @Override
    public short[] getShortArr(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return new short[]{2,6,7};
    }

    @Override
    public short[] getShortArrFrom(short[] arr, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return arr;
    }

    @Override
    public int getInt(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        Random r = new Random();

        return r.nextInt();
    }

    @Override
    public int getIntFrom(int value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx, value);
        return value + 500;
    }

    @Override
    public int[] getIntArr(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        Random r = new Random();
        int size = r.nextInt(20);
        int[] arr = new int[size];
        arr[0] = r.nextInt(10000);
        for (int i = 1; i < size; i++) {
            arr[i] = arr[i-1] + 1;
        }
        return arr;
    }

    @Override
    public int[] getIntArrFrom(int[] arr, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return arr;
    }

    @Override
    public long getLong(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return 922337203685477580L;
    }

    @Override
    public long getLongFrom(long value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},from = {}", method, ctx, value);
        return 92233720368547758L;
    }

    @Override
    public long[] getLongArr(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        Random r = new Random();
        int size = r.nextInt(300);
        long[] arr = new long[size];
        arr[0] = r.nextLong();
        for (int i = 1; i < size; i++) {
            arr[i] = arr[i-1]+1;
        }
        return arr;
    }

    @Override
    public long[] getLongArrFrom(long[] arr, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return arr;
    }

    @Override
    public float getFloat(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx);
        return 442.43f;
    }

    @Override
    public float getFloatFrom(float value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx, value);
        return value += 5;
    }

    @Override
    public float[] getFloatArr(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        Random r = new Random();
        int size = r.nextInt(300);
        float[] arr = new float[size];
        arr[0] = r.nextInt(400);
        for (int i = 1; i < size; i++) {
            arr[i] = arr[i-1]+1;
        }
        return arr;
    }

    @Override
    public float[] getFloatArrFrom(float[] arr, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return arr;
    }

    @Override
    public double getDouble(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={}", method, ctx);
        return 4455.32D;
    }

    @Override
    public double getDoubleFrom(double value, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx, value);
        return value += 100;
    }

    @Override
    public double[] getDoubleArr(Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx);
        Random r = new Random();
        int size = r.nextInt(300);
        double[] arr = new double[size];
        for (int i = 0; i < size; i++) {
            arr[i] = (double) i;
        }
        return arr;
    }

    @Override
    public double[] getDoubleArrFrom(double[] arr, Current current) {
        Map<String, String> ctx = current.ctx;
        String method = Thread.currentThread().getStackTrace()[1].getMethodName();
        logger.info("方法[{}],ctx={},value = {}", method, ctx);

        return arr;
    }

    @Override
    public String execute(Request request, Current current) {
        Map<String, String> ctx = current.ctx;
        String m = Thread.currentThread().getStackTrace()[1].getMethodName();
        String method = request.method;
        Map<String, String> params = request.params;
        HashMap<String, Object> resultMap = new HashMap<>();
        logger.info("方法[{}],ctx={},method = {} ,iceMethod = {},params = {}", ctx, m, method, params);
        if (method.equals("getArticle")) {
            Article article = new Article();

            article.setItem(params.get("item"));
            article.setId(Integer.parseInt(params.get("id")));
            article.setTitle("this is a title!");
            article.setContent("this is a content!");
            resultMap.put("code", 1);
            resultMap.put("message", null);
            resultMap.put("data", article);
            String jsonR = JsonUtil.toJson(resultMap);
            logger.info(jsonR);
            return jsonR;
        } else if (method.equals("exception")) {
            String type = request.params.get("type");
            if (type == null || type.length() == 0) {
                throw new RuntimeException("has exception");
            }
            switch (type) {
                case "zero":
                    int i = 0;
                    i = 122 / i;
                case "timeout":
                    logger.info("超时测试");
                    try {
                        TimeUnit.SECONDS.sleep(20);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
            }
        } else if (method.equals("getLargeString")) {
            StringBuffer buffer = new StringBuffer();
            for (int i = 0; i < 8; i++) {
                buffer.append("aaabbbcccdddeeefffggghhhiiijjj");
            }
            if (params.get("a").equals("a")) {
                buffer.append("0123456789");
                buffer.append("0123456789");
            }
            logger.info(buffer.toString());
            return buffer.toString();
        } else if (method.equals("getStringArticle")) {
            Article article = new Article();
            article.setItem(params.get("item"));
            article.setId(Integer.parseInt(params.get("id")));
            article.setTitle("this is a title!");
            article.setContent("this is a content!");
            resultMap.put("code", "1");
            resultMap.put("message", "");
            HashMap<String, String> rmap = new HashMap<>();
            rmap.put("item", article.getItem());
            rmap.put("id", String.valueOf(article.getId()));
            rmap.put("title", article.getTitle());
            rmap.put("content", article.getContent());
            resultMap.put("data", rmap);
            String jsonR = JsonUtil.toJson(resultMap);
            logger.info(jsonR);
            return jsonR;
        }
        return null;
    }


}
