#include <bits/stdc++.h>
using namespace std;

set< vector<int> > vis;

struct Move {
    vector<int> now;
    int id, dire; // dire 0 ��ʾ���ϻ����� 1 ��ʾ���»�����
};

map< vector<int>, Move > path; 

vector<int> s; // ��ʼ״̬ 6 * 6 �����ƽ��

queue< vector<int> > q;

int dire[20]; // 0 ����ˮƽ���� 1 ����ֱ����

int n, g; // ��ʾ���������gΪĿ���

int e(int x, int y) {
    return x * 6 + y;
}

bool isend(vector<int> & map) { // ����Ƿ��յ�  
    int x, y;
    for(int i=0; i < 6; i++) 
        for(int j=0; j < 6; j++) 
            if(map[e(i,j)] == g) {
                x = i; y = j;
            }
    if(dire[g]) {
        return false;
    }else {
        for(int i=y; i < 6; i++)
            if(map[e(x,i)] != 0 && map[e(x,i)] != g) return false;
    }
    return true;
}

bool check(int x, int y) { // ��飨x, y���Ƿ�Ϸ�
    if(x < 0 || x >= 6 || y < 0 || y >= 6) return false;
    return true;
}

vector<int> move(vector<int> map, int id, int dir) { // ����ͼmap�б��Ϊid�Ŀ鰴dir�����ƶ�
    int x, y;
    for(int i=0; i < 6; i++) 
        for(int j=0; j < 6; j++) 
            if(map[e(i,j)] == id) {
                x = i; y = j;
            } 
    if(dire[id]) {
        if(dir) {
            while(check(x+1, y) && map[e(x+1, y)] == id) x++;
            int nx = x, ny = y;
            while(check(nx+1, ny) && map[e(nx+1, ny)] == 0) nx++;
            while(check(x, y) && map[e(x, y)] == id) {
                map[e(x, y)] = 0;
                map[e(nx, ny)] = id;
                x--; nx--;
            }
        }else {
            while(check(x-1, y) && map[e(x-1, y)] == id) x--;
            int nx = x, ny = y;
            while(check(nx-1, ny) && map[e(nx-1, ny)] == 0) nx--;
            while(check(x, y) && map[e(x, y)] == id) {
                map[e(x, y)] = 0;
                map[e(nx, ny)] = id;
                x++; nx++;
            }
        }
    }else {
        if(dir) {
            while(check(x, y+1) && map[e(x, y+1)] == id) y++;
            int nx = x, ny = y;
            while(check(nx, ny+1) && map[e(nx, ny+1)] == 0) ny++;
            while(check(x, y) && map[e(x, y)] == id) {
                map[e(x, y)] = 0;
                map[e(nx, ny)] = id;
                y--; ny--;
            }
        }else {
            while(check(x, y-1) && map[e(x, y-1)] == id) y--;
            int nx = x, ny = y;
            while(check(nx, ny-1) && map[e(nx, ny-1)] == 0) ny--;
            while(check(x, y) && map[e(x, y)] == id) {
                map[e(x, y)] = 0;
                map[e(nx, ny)] = id;
                y++; ny++;
            }
        }
    }
    return map;
} 


vector<int> bfs() {
    vector<int> now, nxt;
    q.push(s);
    vis.insert(s);
    if(isend(s)) return s;
    while(!q.empty()) {
        now = q.front();
        q.pop();
        for(int i=1; i <= n; i++) {
            nxt = move(now, i, 0);
            if(vis.find(nxt) == vis.end()) {
                vis.insert(nxt);
                path[nxt] = {now, i, 0};
                if(isend(nxt)) return nxt;
                q.push(nxt);
            }
            nxt = move(now, i, 1);
            if(vis.find(nxt) == vis.end()) {
                vis.insert(nxt);
                path[nxt] = {now, i, 1};
                if(isend(nxt)) return nxt;
                q.push(nxt);
            }
        }
    }
    return s;
}

void printmap(vector<int> & map) {
    for(int i=0; i < 6; i++) {
        for(int j=0; j < 6; j++) {
            printf("%d ", map[e(i, j)]);
        }
        putchar('\n');
    }
}

void printres(vector<int> & now) {
    if(now == s) {
        printmap(now);
        return ;
    }
    Move m = path[now];
    printres(m.now);
    putchar('\n');
    printf("%d ", m.id);
    if(dire[m.id]) {
        if(m.dire) {
            printf("�����ƶ�\n");
        }else {
            printf("�����ƶ�\n");
        }
    }else {
        if(m.dire) {
            printf("�����ƶ�\n");
        }else {
            printf("�����ƶ�\n");
        }
    }
    putchar('\n');
    printmap(now);
} 

int main() {
    freopen("in.txt", "r", stdin);
    freopen("out.txt", "w", stdout);
    for(int i=0; i < 6; i++) {
        for(int j=0; j < 6; j++) { 
            int x;
            scanf("%d", &x);
            s.push_back(x);
        }
    }
    scanf("%d%d", &n, &g);
    int d[2][4] = {1,-1,0,0,0,0,1,-1};
    for(int i=0; i < 6; i++) {
        for(int j=0; j < 6; j++) {
            int now = s[e(i, j)];
            for(int k=0; k < 4; k++) {
                int x = i + d[0][k];
                int y = j + d[1][k];
                if(!check(x, y)) continue;
                if(s[e(x, y)] == now) {
                    if(k <= 1) dire[now] = 1;
                    else dire[now] = 0;
                }
            }
        }
    }
    vector<int> myend = bfs();
    // printmap(myend);
    printres(myend);
    return 0;
}
/*

0 0 1 2 2 2
3 0 1 4 0 5
3 6 6 4 0 5
0 7 0 4 0 8
0 7 9 9 0 8
0 0 0 0 0 0
9 6

*/