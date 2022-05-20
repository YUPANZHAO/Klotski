#include "TaskQueue.h"

template <class T>
TaskQueue<T>::TaskQueue() {
    pthread_mutex_init(&m_mutex, NULL);
}

template <class T>
TaskQueue<T>::~TaskQueue() {
    pthread_mutex_destroy(&m_mutex);
}

template <class T>
void TaskQueue<T>::addTask(Task<T> task) {
    pthread_mutex_lock(&m_mutex);
    m_taskQ.push(task);
    pthread_mutex_unlock(&m_mutex);
}

template <class T>
void TaskQueue<T>::addTask(callback f, void* arg) {
    pthread_mutex_lock(&m_mutex);
    m_taskQ.push(Task<T>(f, arg));
    pthread_mutex_unlock(&m_mutex);
}

template <class T>
Task<T> TaskQueue<T>::takeTask() {
    Task<T> t;
    pthread_mutex_lock(&m_mutex);
    if(!m_taskQ.empty()) {
        t = m_taskQ.front();
        m_taskQ.pop();
    }
    pthread_mutex_unlock(&m_mutex);
    return t;
}