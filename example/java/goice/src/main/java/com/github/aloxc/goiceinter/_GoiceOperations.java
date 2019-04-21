// **********************************************************************
//
// Copyright (c) 2003-2017 ZeroC, Inc. All rights reserved.
//
// This copy of Ice is licensed to you under the terms described in the
// ICE_LICENSE file included in this distribution.
//
// **********************************************************************
//
// Ice version 3.6.4
//
// <auto-generated>
//
// Generated from file `goice.ice'
//
// Warning: do not edit this file.
//
// </auto-generated>
//

package com.github.aloxc.goiceinter;

public interface _GoiceOperations
{
    String non(Ice.Current __current);

    String getString(Ice.Current __current);

    String getStringFrom(String value, Ice.Current __current);

    String[] getStringArr(Ice.Current __current);

    String[] getStringArrFrom(String[] arr, Ice.Current __current);

    String two(String from, String to, Ice.Current __current);

    void vvoid(Ice.Current __current);

    void vvoidTo(String value, Ice.Current __current);

    byte getByte(Ice.Current __current);

    byte getByteFrom(byte value, Ice.Current __current);

    byte[] getByteArr(Ice.Current __current);

    byte[] getByteArrFrom(byte[] arr, Ice.Current __current);

    boolean getBool(Ice.Current __current);

    boolean getBoolFrom(boolean value, Ice.Current __current);

    boolean[] getBoolArr(Ice.Current __current);

    boolean[] getBoolArrFrom(boolean[] arr, Ice.Current __current);

    short getShort(Ice.Current __current);

    short getShortFrom(short value, Ice.Current __current);

    short[] getShortArr(Ice.Current __current);

    short[] getShortArrFrom(short[] arr, Ice.Current __current);

    int getInt(Ice.Current __current);

    int getIntFrom(int value, Ice.Current __current);

    int[] getIntArr(Ice.Current __current);

    int[] getIntArrFrom(int[] arr, Ice.Current __current);

    long getLong(Ice.Current __current);

    long getLongFrom(long value, Ice.Current __current);

    long[] getLongArr(Ice.Current __current);

    long[] getLongArrFrom(long[] arr, Ice.Current __current);

    float getFloat(Ice.Current __current);

    float getFloatFrom(float value, Ice.Current __current);

    float[] getFloatArr(Ice.Current __current);

    float[] getFloatArrFrom(float[] arr, Ice.Current __current);

    double getDouble(Ice.Current __current);

    double getDoubleFrom(double value, Ice.Current __current);

    double[] getDoubleArr(Ice.Current __current);

    double[] getDoubleArrFrom(double[] arr, Ice.Current __current);

    String execute(Request request, Ice.Current __current);
}
