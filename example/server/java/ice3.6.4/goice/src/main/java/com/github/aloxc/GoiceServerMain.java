package com.github.aloxc;

import Ice.Communicator;
import Ice.Identity;
import Ice.ObjectAdapter;
import Ice.Util;

public class GoiceServerMain {
    public static void main(String[] args) {
        // 通信器
        Communicator ic = null;
        // 初始化这个通信器
        ic = Util.initialize(args);
        // 创建ice适配器,将服务调用地址和服务映射起来
        // "HelloServiceAdapter"是适配器名, "default -p 1888"是服务调用的地址
//        ObjectAdapter adapter = ic.createObjectAdapterWithEndpoints("HelloServiceAdapter","default -p 1888");
        ObjectAdapter adapter = ic.createObjectAdapterWithEndpoints("Goicedapter","default -p 1888");
        // 将服务的具体实现类servant交给这个适配器
        Ice.Object servant = new GoiceImpl();
        // "HelloIce"--服务接口在ice中注册名,转成唯一标识identity
        Identity id = Util.stringToIdentity("Goice");
        System.out.println("name = "+id.name);
        System.out.println("adapter = " +adapter.getName());
        System.out.println("id = " + servant.ice_id());
        adapter.add(servant, id);
        // 激活这个适配器
        adapter.activate();

        System.out.println("server服务容器启动成功。。。");
    }
}
