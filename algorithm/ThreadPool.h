#pragma once
#include "TaskQueue.h"
#include "TaskQueue.cpp"

template <class T>
class ThreadPool {
public:
    //创建线程池
    ThreadPool(int min, int max);
    //销毁线程池
    ~ThreadPool();
    //给线程池添加任务
    void addTask(Task<T> task);
    //获取线程池中工作的线程个数
    int getBusyNum();
    //获取线程池中活着的线程个数
    int getAliveNum();

private:
    //工作者线程
    static void* worker(void* arg);
    //管理者线程
    static void* manager(void* arg);
    //单个线程退出函数
    void threadExit();
    
private:
    TaskQueue<T>* taskQ;       //任务队列
    
    pthread_t managerID;    //管理者线程ID
    pthread_t* threadIDs;   //工作线程ID
    
    int minNum;     //最小线程数
    int maxNum;     //最大线程数
    int busyNum;    //忙的线程数
    int liveNum;    //存活线程数
    int exitNum;    //要销毁的线程数

    pthread_mutex_t mutexPool;  //锁整个线程池
    pthread_cond_t notEmpty;    //任务队列是不是空了
    
    bool shutdown;   //是不是要销毁线程池，销毁为1，不销毁为0

    static const int NUMBER = 2;
};