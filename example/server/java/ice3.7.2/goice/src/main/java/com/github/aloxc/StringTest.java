package com.github.aloxc;

public class StringTest {
    public static void main(String[] args) {
        String s = "a";
        StringBuffer buffer = new StringBuffer();
        int[] arr = new int[50];
        for (int i = 0; i < 50; i++) {
            buffer.append(i);
            s = "a"+i;
            arr[i] = i;
            System.out.println(s.hashCode() + " " + arr+" " + buffer.hashCode());
        }
    }
}
