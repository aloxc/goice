package com.github.aloxc;

import com.zeroc.Ice.*;
import com.github.aloxc.goiceinter.GoicePrx;
import com.github.aloxc.user.post.UserpostPrx;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.lang.Object;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;

public class IceServerMain {
    public final static Logger logger = LoggerFactory.getLogger(IceServerMain.class);
    private static Map<String, com.zeroc.Ice.Object> servantCache = new ConcurrentHashMap<String, com.zeroc.Ice.Object>(200, 0.75f, 256);

    public static void main(String[] args) {
        InitializationData initData = new InitializationData();
        Properties prop = Util.createProperties();
        initData.properties = prop;

        Communicator ic = Util.initialize(initData);
        String adapterId = "Goice";
        String endpoint = "default -p 1888";
//        ObjectAdapter adapter = ic.createObjectAdapter("Goice:default -p 1888");
        ObjectAdapter adapter = ic.createObjectAdapterWithEndpoints(adapterId, endpoint);
        adapter.addServantLocator(new ServantLocator() {
            @Override
            public LocateResult locate(Current current) throws UserException {
                return null;
            }

            @Override
            public void finished(Current current, com.zeroc.Ice.Object object, Object o) throws UserException {

            }

            @Override
            public void deactivate(String s) {

            }
        },"");

        adapter.activate();
    }

    public static com.zeroc.Ice.Object getServantBean(String strToProxy) {

        com.zeroc.Ice.Object servant = servantCache.get(strToProxy);
        if (null == servant) {
            String itfc = null;
            String impl = null;
            if (strToProxy.equals("Goice")) {
                itfc = GoicePrx.class.getName();
                impl = GoiceImpl.class.getName();
            } else if (strToProxy.equals("Userpost")) {
                itfc = UserpostPrx.class.getName();
                impl = UserpostImpl.class.getName();
            }
            servant = loadServantInstance(itfc, impl);
            servantCache.put(strToProxy, servant);
        }
        return servant;
    }

    private static com.zeroc.Ice.Object loadServantInstance(final String itfc, final String impl) {

        try {
            Class clazz = Class.forName(impl);
            Object target = clazz.newInstance();
            return (com.zeroc.Ice.Object) target;
        } catch (Throwable thr) {
            throw new RuntimeException(thr.getCause());
        }
    }
}
