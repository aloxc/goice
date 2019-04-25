//
// Copyright (c) ZeroC, Inc. All rights reserved.
//
//
// Ice version 3.7.2
//
// <auto-generated>
//
// Generated from file `userpost.ice'
//
// Warning: do not edit this file.
//
// </auto-generated>
//

package com.github.aloxc.user.post;

public interface UserpostPrx extends com.zeroc.Ice.ObjectPrx
{
    default String valuestr(String name, String value)
    {
        return valuestr(name, value, com.zeroc.Ice.ObjectPrx.noExplicitContext);
    }

    default String valuestr(String name, String value, java.util.Map<String, String> context)
    {
        return _iceI_valuestrAsync(name, value, context, true).waitForResponse();
    }

    default java.util.concurrent.CompletableFuture<String> valuestrAsync(String name, String value)
    {
        return _iceI_valuestrAsync(name, value, com.zeroc.Ice.ObjectPrx.noExplicitContext, false);
    }

    default java.util.concurrent.CompletableFuture<String> valuestrAsync(String name, String value, java.util.Map<String, String> context)
    {
        return _iceI_valuestrAsync(name, value, context, false);
    }

    /**
     * @hidden
     * @param iceP_name -
     * @param iceP_value -
     * @param context -
     * @param sync -
     * @return -
     **/
    default com.zeroc.IceInternal.OutgoingAsync<String> _iceI_valuestrAsync(String iceP_name, String iceP_value, java.util.Map<String, String> context, boolean sync)
    {
        com.zeroc.IceInternal.OutgoingAsync<String> f = new com.zeroc.IceInternal.OutgoingAsync<>(this, "valuestr", null, sync, null);
        f.invoke(true, context, null, ostr -> {
                     ostr.writeString(iceP_name);
                     ostr.writeString(iceP_value);
                 }, istr -> {
                     String ret;
                     ret = istr.readString();
                     return ret;
                 });
        return f;
    }

    default String valuelong(String name, long value)
    {
        return valuelong(name, value, com.zeroc.Ice.ObjectPrx.noExplicitContext);
    }

    default String valuelong(String name, long value, java.util.Map<String, String> context)
    {
        return _iceI_valuelongAsync(name, value, context, true).waitForResponse();
    }

    default java.util.concurrent.CompletableFuture<String> valuelongAsync(String name, long value)
    {
        return _iceI_valuelongAsync(name, value, com.zeroc.Ice.ObjectPrx.noExplicitContext, false);
    }

    default java.util.concurrent.CompletableFuture<String> valuelongAsync(String name, long value, java.util.Map<String, String> context)
    {
        return _iceI_valuelongAsync(name, value, context, false);
    }

    /**
     * @hidden
     * @param iceP_name -
     * @param iceP_value -
     * @param context -
     * @param sync -
     * @return -
     **/
    default com.zeroc.IceInternal.OutgoingAsync<String> _iceI_valuelongAsync(String iceP_name, long iceP_value, java.util.Map<String, String> context, boolean sync)
    {
        com.zeroc.IceInternal.OutgoingAsync<String> f = new com.zeroc.IceInternal.OutgoingAsync<>(this, "valuelong", null, sync, null);
        f.invoke(true, context, null, ostr -> {
                     ostr.writeString(iceP_name);
                     ostr.writeLong(iceP_value);
                 }, istr -> {
                     String ret;
                     ret = istr.readString();
                     return ret;
                 });
        return f;
    }

    default String threeparams(String name, int value, double any)
    {
        return threeparams(name, value, any, com.zeroc.Ice.ObjectPrx.noExplicitContext);
    }

    default String threeparams(String name, int value, double any, java.util.Map<String, String> context)
    {
        return _iceI_threeparamsAsync(name, value, any, context, true).waitForResponse();
    }

    default java.util.concurrent.CompletableFuture<String> threeparamsAsync(String name, int value, double any)
    {
        return _iceI_threeparamsAsync(name, value, any, com.zeroc.Ice.ObjectPrx.noExplicitContext, false);
    }

    default java.util.concurrent.CompletableFuture<String> threeparamsAsync(String name, int value, double any, java.util.Map<String, String> context)
    {
        return _iceI_threeparamsAsync(name, value, any, context, false);
    }

    /**
     * @hidden
     * @param iceP_name -
     * @param iceP_value -
     * @param iceP_any -
     * @param context -
     * @param sync -
     * @return -
     **/
    default com.zeroc.IceInternal.OutgoingAsync<String> _iceI_threeparamsAsync(String iceP_name, int iceP_value, double iceP_any, java.util.Map<String, String> context, boolean sync)
    {
        com.zeroc.IceInternal.OutgoingAsync<String> f = new com.zeroc.IceInternal.OutgoingAsync<>(this, "threeparams", null, sync, null);
        f.invoke(true, context, null, ostr -> {
                     ostr.writeString(iceP_name);
                     ostr.writeInt(iceP_value);
                     ostr.writeDouble(iceP_any);
                 }, istr -> {
                     String ret;
                     ret = istr.readString();
                     return ret;
                 });
        return f;
    }

    default String doarr(String[] sar, int[] iarr, String key, String value, int i)
    {
        return doarr(sar, iarr, key, value, i, com.zeroc.Ice.ObjectPrx.noExplicitContext);
    }

    default String doarr(String[] sar, int[] iarr, String key, String value, int i, java.util.Map<String, String> context)
    {
        return _iceI_doarrAsync(sar, iarr, key, value, i, context, true).waitForResponse();
    }

    default java.util.concurrent.CompletableFuture<String> doarrAsync(String[] sar, int[] iarr, String key, String value, int i)
    {
        return _iceI_doarrAsync(sar, iarr, key, value, i, com.zeroc.Ice.ObjectPrx.noExplicitContext, false);
    }

    default java.util.concurrent.CompletableFuture<String> doarrAsync(String[] sar, int[] iarr, String key, String value, int i, java.util.Map<String, String> context)
    {
        return _iceI_doarrAsync(sar, iarr, key, value, i, context, false);
    }

    /**
     * @hidden
     * @param iceP_sar -
     * @param iceP_iarr -
     * @param iceP_key -
     * @param iceP_value -
     * @param iceP_i -
     * @param context -
     * @param sync -
     * @return -
     **/
    default com.zeroc.IceInternal.OutgoingAsync<String> _iceI_doarrAsync(String[] iceP_sar, int[] iceP_iarr, String iceP_key, String iceP_value, int iceP_i, java.util.Map<String, String> context, boolean sync)
    {
        com.zeroc.IceInternal.OutgoingAsync<String> f = new com.zeroc.IceInternal.OutgoingAsync<>(this, "doarr", null, sync, null);
        f.invoke(true, context, null, ostr -> {
                     ostr.writeStringSeq(iceP_sar);
                     ostr.writeIntSeq(iceP_iarr);
                     ostr.writeString(iceP_key);
                     ostr.writeString(iceP_value);
                     ostr.writeInt(iceP_i);
                 }, istr -> {
                     String ret;
                     ret = istr.readString();
                     return ret;
                 });
        return f;
    }

    default String todo(String json)
    {
        return todo(json, com.zeroc.Ice.ObjectPrx.noExplicitContext);
    }

    default String todo(String json, java.util.Map<String, String> context)
    {
        return _iceI_todoAsync(json, context, true).waitForResponse();
    }

    default java.util.concurrent.CompletableFuture<String> todoAsync(String json)
    {
        return _iceI_todoAsync(json, com.zeroc.Ice.ObjectPrx.noExplicitContext, false);
    }

    default java.util.concurrent.CompletableFuture<String> todoAsync(String json, java.util.Map<String, String> context)
    {
        return _iceI_todoAsync(json, context, false);
    }

    /**
     * @hidden
     * @param iceP_json -
     * @param context -
     * @param sync -
     * @return -
     **/
    default com.zeroc.IceInternal.OutgoingAsync<String> _iceI_todoAsync(String iceP_json, java.util.Map<String, String> context, boolean sync)
    {
        com.zeroc.IceInternal.OutgoingAsync<String> f = new com.zeroc.IceInternal.OutgoingAsync<>(this, "todo", null, sync, null);
        f.invoke(true, context, null, ostr -> {
                     ostr.writeString(iceP_json);
                 }, istr -> {
                     String ret;
                     ret = istr.readString();
                     return ret;
                 });
        return f;
    }

    default int[] getIntArr(int i)
    {
        return getIntArr(i, com.zeroc.Ice.ObjectPrx.noExplicitContext);
    }

    default int[] getIntArr(int i, java.util.Map<String, String> context)
    {
        return _iceI_getIntArrAsync(i, context, true).waitForResponse();
    }

    default java.util.concurrent.CompletableFuture<int[]> getIntArrAsync(int i)
    {
        return _iceI_getIntArrAsync(i, com.zeroc.Ice.ObjectPrx.noExplicitContext, false);
    }

    default java.util.concurrent.CompletableFuture<int[]> getIntArrAsync(int i, java.util.Map<String, String> context)
    {
        return _iceI_getIntArrAsync(i, context, false);
    }

    /**
     * @hidden
     * @param iceP_i -
     * @param context -
     * @param sync -
     * @return -
     **/
    default com.zeroc.IceInternal.OutgoingAsync<int[]> _iceI_getIntArrAsync(int iceP_i, java.util.Map<String, String> context, boolean sync)
    {
        com.zeroc.IceInternal.OutgoingAsync<int[]> f = new com.zeroc.IceInternal.OutgoingAsync<>(this, "getIntArr", null, sync, null);
        f.invoke(true, context, null, ostr -> {
                     ostr.writeInt(iceP_i);
                 }, istr -> {
                     int[] ret;
                     ret = istr.readIntSeq();
                     return ret;
                 });
        return f;
    }

    default String[] getStrArr(int i)
    {
        return getStrArr(i, com.zeroc.Ice.ObjectPrx.noExplicitContext);
    }

    default String[] getStrArr(int i, java.util.Map<String, String> context)
    {
        return _iceI_getStrArrAsync(i, context, true).waitForResponse();
    }

    default java.util.concurrent.CompletableFuture<String[]> getStrArrAsync(int i)
    {
        return _iceI_getStrArrAsync(i, com.zeroc.Ice.ObjectPrx.noExplicitContext, false);
    }

    default java.util.concurrent.CompletableFuture<String[]> getStrArrAsync(int i, java.util.Map<String, String> context)
    {
        return _iceI_getStrArrAsync(i, context, false);
    }

    /**
     * @hidden
     * @param iceP_i -
     * @param context -
     * @param sync -
     * @return -
     **/
    default com.zeroc.IceInternal.OutgoingAsync<String[]> _iceI_getStrArrAsync(int iceP_i, java.util.Map<String, String> context, boolean sync)
    {
        com.zeroc.IceInternal.OutgoingAsync<String[]> f = new com.zeroc.IceInternal.OutgoingAsync<>(this, "getStrArr", null, sync, null);
        f.invoke(true, context, null, ostr -> {
                     ostr.writeInt(iceP_i);
                 }, istr -> {
                     String[] ret;
                     ret = istr.readStringSeq();
                     return ret;
                 });
        return f;
    }

    /**
     * Contacts the remote server to verify that the object implements this type.
     * Raises a local exception if a communication error occurs.
     * @param obj The untyped proxy.
     * @return A proxy for this type, or null if the object does not support this type.
     **/
    static UserpostPrx checkedCast(com.zeroc.Ice.ObjectPrx obj)
    {
        return com.zeroc.Ice.ObjectPrx._checkedCast(obj, ice_staticId(), UserpostPrx.class, _UserpostPrxI.class);
    }

    /**
     * Contacts the remote server to verify that the object implements this type.
     * Raises a local exception if a communication error occurs.
     * @param obj The untyped proxy.
     * @param context The Context map to send with the invocation.
     * @return A proxy for this type, or null if the object does not support this type.
     **/
    static UserpostPrx checkedCast(com.zeroc.Ice.ObjectPrx obj, java.util.Map<String, String> context)
    {
        return com.zeroc.Ice.ObjectPrx._checkedCast(obj, context, ice_staticId(), UserpostPrx.class, _UserpostPrxI.class);
    }

    /**
     * Contacts the remote server to verify that a facet of the object implements this type.
     * Raises a local exception if a communication error occurs.
     * @param obj The untyped proxy.
     * @param facet The name of the desired facet.
     * @return A proxy for this type, or null if the object does not support this type.
     **/
    static UserpostPrx checkedCast(com.zeroc.Ice.ObjectPrx obj, String facet)
    {
        return com.zeroc.Ice.ObjectPrx._checkedCast(obj, facet, ice_staticId(), UserpostPrx.class, _UserpostPrxI.class);
    }

    /**
     * Contacts the remote server to verify that a facet of the object implements this type.
     * Raises a local exception if a communication error occurs.
     * @param obj The untyped proxy.
     * @param facet The name of the desired facet.
     * @param context The Context map to send with the invocation.
     * @return A proxy for this type, or null if the object does not support this type.
     **/
    static UserpostPrx checkedCast(com.zeroc.Ice.ObjectPrx obj, String facet, java.util.Map<String, String> context)
    {
        return com.zeroc.Ice.ObjectPrx._checkedCast(obj, facet, context, ice_staticId(), UserpostPrx.class, _UserpostPrxI.class);
    }

    /**
     * Downcasts the given proxy to this type without contacting the remote server.
     * @param obj The untyped proxy.
     * @return A proxy for this type.
     **/
    static UserpostPrx uncheckedCast(com.zeroc.Ice.ObjectPrx obj)
    {
        return com.zeroc.Ice.ObjectPrx._uncheckedCast(obj, UserpostPrx.class, _UserpostPrxI.class);
    }

    /**
     * Downcasts the given proxy to this type without contacting the remote server.
     * @param obj The untyped proxy.
     * @param facet The name of the desired facet.
     * @return A proxy for this type.
     **/
    static UserpostPrx uncheckedCast(com.zeroc.Ice.ObjectPrx obj, String facet)
    {
        return com.zeroc.Ice.ObjectPrx._uncheckedCast(obj, facet, UserpostPrx.class, _UserpostPrxI.class);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the per-proxy context.
     * @param newContext The context for the new proxy.
     * @return A proxy with the specified per-proxy context.
     **/
    @Override
    default UserpostPrx ice_context(java.util.Map<String, String> newContext)
    {
        return (UserpostPrx)_ice_context(newContext);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the adapter ID.
     * @param newAdapterId The adapter ID for the new proxy.
     * @return A proxy with the specified adapter ID.
     **/
    @Override
    default UserpostPrx ice_adapterId(String newAdapterId)
    {
        return (UserpostPrx)_ice_adapterId(newAdapterId);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the endpoints.
     * @param newEndpoints The endpoints for the new proxy.
     * @return A proxy with the specified endpoints.
     **/
    @Override
    default UserpostPrx ice_endpoints(com.zeroc.Ice.Endpoint[] newEndpoints)
    {
        return (UserpostPrx)_ice_endpoints(newEndpoints);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the locator cache timeout.
     * @param newTimeout The new locator cache timeout (in seconds).
     * @return A proxy with the specified locator cache timeout.
     **/
    @Override
    default UserpostPrx ice_locatorCacheTimeout(int newTimeout)
    {
        return (UserpostPrx)_ice_locatorCacheTimeout(newTimeout);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the invocation timeout.
     * @param newTimeout The new invocation timeout (in seconds).
     * @return A proxy with the specified invocation timeout.
     **/
    @Override
    default UserpostPrx ice_invocationTimeout(int newTimeout)
    {
        return (UserpostPrx)_ice_invocationTimeout(newTimeout);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for connection caching.
     * @param newCache <code>true</code> if the new proxy should cache connections; <code>false</code> otherwise.
     * @return A proxy with the specified caching policy.
     **/
    @Override
    default UserpostPrx ice_connectionCached(boolean newCache)
    {
        return (UserpostPrx)_ice_connectionCached(newCache);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the endpoint selection policy.
     * @param newType The new endpoint selection policy.
     * @return A proxy with the specified endpoint selection policy.
     **/
    @Override
    default UserpostPrx ice_endpointSelection(com.zeroc.Ice.EndpointSelectionType newType)
    {
        return (UserpostPrx)_ice_endpointSelection(newType);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for how it selects endpoints.
     * @param b If <code>b</code> is <code>true</code>, only endpoints that use a secure transport are
     * used by the new proxy. If <code>b</code> is false, the returned proxy uses both secure and
     * insecure endpoints.
     * @return A proxy with the specified selection policy.
     **/
    @Override
    default UserpostPrx ice_secure(boolean b)
    {
        return (UserpostPrx)_ice_secure(b);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the encoding used to marshal parameters.
     * @param e The encoding version to use to marshal request parameters.
     * @return A proxy with the specified encoding version.
     **/
    @Override
    default UserpostPrx ice_encodingVersion(com.zeroc.Ice.EncodingVersion e)
    {
        return (UserpostPrx)_ice_encodingVersion(e);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for its endpoint selection policy.
     * @param b If <code>b</code> is <code>true</code>, the new proxy will use secure endpoints for invocations
     * and only use insecure endpoints if an invocation cannot be made via secure endpoints. If <code>b</code> is
     * <code>false</code>, the proxy prefers insecure endpoints to secure ones.
     * @return A proxy with the specified selection policy.
     **/
    @Override
    default UserpostPrx ice_preferSecure(boolean b)
    {
        return (UserpostPrx)_ice_preferSecure(b);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the router.
     * @param router The router for the new proxy.
     * @return A proxy with the specified router.
     **/
    @Override
    default UserpostPrx ice_router(com.zeroc.Ice.RouterPrx router)
    {
        return (UserpostPrx)_ice_router(router);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for the locator.
     * @param locator The locator for the new proxy.
     * @return A proxy with the specified locator.
     **/
    @Override
    default UserpostPrx ice_locator(com.zeroc.Ice.LocatorPrx locator)
    {
        return (UserpostPrx)_ice_locator(locator);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for collocation optimization.
     * @param b <code>true</code> if the new proxy enables collocation optimization; <code>false</code> otherwise.
     * @return A proxy with the specified collocation optimization.
     **/
    @Override
    default UserpostPrx ice_collocationOptimized(boolean b)
    {
        return (UserpostPrx)_ice_collocationOptimized(b);
    }

    /**
     * Returns a proxy that is identical to this proxy, but uses twoway invocations.
     * @return A proxy that uses twoway invocations.
     **/
    @Override
    default UserpostPrx ice_twoway()
    {
        return (UserpostPrx)_ice_twoway();
    }

    /**
     * Returns a proxy that is identical to this proxy, but uses oneway invocations.
     * @return A proxy that uses oneway invocations.
     **/
    @Override
    default UserpostPrx ice_oneway()
    {
        return (UserpostPrx)_ice_oneway();
    }

    /**
     * Returns a proxy that is identical to this proxy, but uses batch oneway invocations.
     * @return A proxy that uses batch oneway invocations.
     **/
    @Override
    default UserpostPrx ice_batchOneway()
    {
        return (UserpostPrx)_ice_batchOneway();
    }

    /**
     * Returns a proxy that is identical to this proxy, but uses datagram invocations.
     * @return A proxy that uses datagram invocations.
     **/
    @Override
    default UserpostPrx ice_datagram()
    {
        return (UserpostPrx)_ice_datagram();
    }

    /**
     * Returns a proxy that is identical to this proxy, but uses batch datagram invocations.
     * @return A proxy that uses batch datagram invocations.
     **/
    @Override
    default UserpostPrx ice_batchDatagram()
    {
        return (UserpostPrx)_ice_batchDatagram();
    }

    /**
     * Returns a proxy that is identical to this proxy, except for compression.
     * @param co <code>true</code> enables compression for the new proxy; <code>false</code> disables compression.
     * @return A proxy with the specified compression setting.
     **/
    @Override
    default UserpostPrx ice_compress(boolean co)
    {
        return (UserpostPrx)_ice_compress(co);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for its connection timeout setting.
     * @param t The connection timeout for the proxy in milliseconds.
     * @return A proxy with the specified timeout.
     **/
    @Override
    default UserpostPrx ice_timeout(int t)
    {
        return (UserpostPrx)_ice_timeout(t);
    }

    /**
     * Returns a proxy that is identical to this proxy, except for its connection ID.
     * @param connectionId The connection ID for the new proxy. An empty string removes the connection ID.
     * @return A proxy with the specified connection ID.
     **/
    @Override
    default UserpostPrx ice_connectionId(String connectionId)
    {
        return (UserpostPrx)_ice_connectionId(connectionId);
    }

    /**
     * Returns a proxy that is identical to this proxy, except it's a fixed proxy bound
     * the given connection.@param connection The fixed proxy connection.
     * @return A fixed proxy bound to the given connection.
     **/
    @Override
    default UserpostPrx ice_fixed(com.zeroc.Ice.Connection connection)
    {
        return (UserpostPrx)_ice_fixed(connection);
    }

    static String ice_staticId()
    {
        return "::user::post::Userpost";
    }
}
