#pragma once
#include <queue>
#include <pthread.h>

using callback = void (*) (void* arg);
//任务结构体
template <class T>
struct Task {
    callback function;
    T* arg;
    Task() {
        function = nullptr;
        arg = nullptr;
    }
    Task(callback f, void* arg) {
        function = f;
        this->arg = (T*)arg;
    }
};
//任务队列
template <class T>
class TaskQueue {
public:
    TaskQueue();
    ~TaskQueue();
    //添加任务
    void addTask(Task<T> task);
    void addTask(callback f, void* arg);
    //取出任务
    Task<T> takeTask();
    //获取当前任务个数
    inline size_t taskNumber() { return m_taskQ.size(); }
private:
    std::queue<Task<T>> m_taskQ;
    pthread_mutex_t m_mutex;
};
