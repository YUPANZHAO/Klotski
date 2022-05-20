#include <iostream>
#include <sys/socket.h>
#include <sys/epoll.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <sys/types.h>
#include <ctype.h>
#include <fcntl.h>
#include "ThreadPool.h"
#include "ThreadPool.cpp"
#include "klotski.cpp"
using namespace std;

//线程参数结构
struct SocketInfo {
    int fd;
    int epfd;
};
//线程池
ThreadPool<SocketInfo> threadPool(10, 100);

void do_accept(void* arg) {
    SocketInfo* info = static_cast<SocketInfo*>(arg);
    int cfd = accept(info->fd, NULL, NULL);
    if(cfd == -1) {
        cerr << "accept connect error" << endl;
        return;
    }
    epoll_event ev;
    ev.events = EPOLLIN | EPOLLET;
    ev.data.fd = cfd;
    int ret = epoll_ctl(info->epfd, EPOLL_CTL_ADD, cfd, &ev);
    if(ret == -1) {
        cerr << "epoll_ctl add error" << endl;
    }
}

void solve(void* arg) {
    SocketInfo* info = static_cast<SocketInfo*>(arg);
    char buf[1024];
    size_t size = sizeof(buf) / sizeof(buf[0]);
    string data = "";
    //读取数据
    while(true) {
        int len = recv(info->fd, buf, size, 0);
        //用户断开连接
        if(len == 0) {
            epoll_ctl(info->epfd, EPOLL_CTL_DEL, info->fd, NULL);
            close(info->fd);
            return;
        }
        if(len == -1) {
            if(errno == EAGAIN) break; //数据读取完毕
            else {
                cerr << "recv data error" << endl;
                return;
            }
        }
        for(int i=0; i < len; ++i) data += buf[i];
    }
    //执行算法
    Klotski algo;
    string res = algo.solve(data);
    //发送数据
    int ret = send(info->fd, res.data(), res.length(), 0);
    if(ret == -1) {
        cerr << "send res error" << endl;
    }
}

int main() {
    //创建socket
    int lfd = socket(AF_INET, SOCK_STREAM, 0);
    if(lfd == -1) {
        cerr << "create socket error" << endl;
        return -1;
    }
    //配置地址
    sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(4331);
    addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    //绑定地址
    int ret = bind(lfd, (sockaddr*)&addr, sizeof(addr));
    if(ret == -1) {
        cerr << "bind address error" << endl;
        return -1;
    }
    //启动监听
    ret = listen(lfd, 64);
    if(ret == -1) {
        cerr << "listen error" << endl;
        return -1;
    }
    //创建epoll
    epoll_event ev;
    ev.events = EPOLLIN | EPOLLET;
    ev.data.fd = lfd;
    int epfd = epoll_create(1);
    epoll_ctl(epfd, EPOLL_CTL_ADD, lfd, &ev);
    //开始处理连接及数据
    while(true) {
        epoll_event evs[1024];
        int size = sizeof(evs) / sizeof(evs[0]);
        int num = epoll_wait(epfd, evs, size, -1);
        if(num == -1) {
            cerr << "epoll wait error" << endl;
            break;
        }
        for(int i=0; i < num; ++i) {
            int fd = evs[i].data.fd;
            int flag = fcntl(fd, F_GETFL);
            flag |= O_NONBLOCK;
            fcntl(fd, F_SETFL, flag);
            //实例化参数
            SocketInfo* info = new SocketInfo;
            info->fd = fd;
            info->epfd = epfd;
            if(fd == lfd) {
                threadPool.addTask(Task<SocketInfo>(do_accept, info));
            }else {
                threadPool.addTask(Task<SocketInfo>(solve, info));
            }
        }
    }
    //断开连接
    close(lfd);
    return 0;
}