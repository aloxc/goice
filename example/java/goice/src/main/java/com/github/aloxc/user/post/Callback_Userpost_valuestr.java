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

public abstract class Callback_Userpost_valuestr
    extends IceInternal.TwowayCallback implements Ice.TwowayCallbackArg1<String>
{
    public final void __completed(Ice.AsyncResult __result)
    {
        UserpostPrxHelper.__valuestr_completed(this, __result);
    }
}
