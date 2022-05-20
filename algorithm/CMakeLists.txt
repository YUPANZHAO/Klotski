cmake_minimum_required (VERSION 3.6.2)

project("KlotskiSolve")

add_definitions("-std=c++11")

add_library(
    TaskQueue
    TaskQueue.h
    TaskQueue.cpp
) 
add_library(
    ThreadPool
    ThreadPool.h
    ThreadPool.cpp
)

add_executable(${PROJECT_NAME} klotski.cpp main.cpp)

target_link_libraries(${PROJECT_NAME} PRIVATE pthread)
target_link_libraries(${PROJECT_NAME} PRIVATE TaskQueue)
target_link_libraries(${PROJECT_NAME} PRIVATE ThreadPool)