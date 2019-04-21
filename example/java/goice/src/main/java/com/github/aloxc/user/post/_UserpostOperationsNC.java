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
// Generated from file `userpost.ice'
//
// Warning: do not edit this file.
//
// </auto-generated>
//

package com.github.aloxc.user.post;

public interface _UserpostOperationsNC
{
    String valuestr(String name, String value);

    String valuelong(String name, long value);

    String threeparams(String name, int value, double any);

    String doarr(String[] sar, int[] iarr, String key, String value, int i);

    String todo(String json);

    int[] getIntArr(int i);

    String[] getStrArr(int i);
}
