# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.6

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/local/bin/cmake

# The command to remove a file.
RM = /usr/local/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/zyp1/Klotski/algorithm

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/zyp1/Klotski/algorithm/build

# Include any dependencies generated for this target.
include CMakeFiles/KlotskiSolve.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/KlotskiSolve.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/KlotskiSolve.dir/flags.make

CMakeFiles/KlotskiSolve.dir/klotski.cpp.o: CMakeFiles/KlotskiSolve.dir/flags.make
CMakeFiles/KlotskiSolve.dir/klotski.cpp.o: ../klotski.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/zyp1/Klotski/algorithm/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/KlotskiSolve.dir/klotski.cpp.o"
	/usr/bin/c++   $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/KlotskiSolve.dir/klotski.cpp.o -c /home/zyp1/Klotski/algorithm/klotski.cpp

CMakeFiles/KlotskiSolve.dir/klotski.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/KlotskiSolve.dir/klotski.cpp.i"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/zyp1/Klotski/algorithm/klotski.cpp > CMakeFiles/KlotskiSolve.dir/klotski.cpp.i

CMakeFiles/KlotskiSolve.dir/klotski.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/KlotskiSolve.dir/klotski.cpp.s"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/zyp1/Klotski/algorithm/klotski.cpp -o CMakeFiles/KlotskiSolve.dir/klotski.cpp.s

CMakeFiles/KlotskiSolve.dir/klotski.cpp.o.requires:

.PHONY : CMakeFiles/KlotskiSolve.dir/klotski.cpp.o.requires

CMakeFiles/KlotskiSolve.dir/klotski.cpp.o.provides: CMakeFiles/KlotskiSolve.dir/klotski.cpp.o.requires
	$(MAKE) -f CMakeFiles/KlotskiSolve.dir/build.make CMakeFiles/KlotskiSolve.dir/klotski.cpp.o.provides.build
.PHONY : CMakeFiles/KlotskiSolve.dir/klotski.cpp.o.provides

CMakeFiles/KlotskiSolve.dir/klotski.cpp.o.provides.build: CMakeFiles/KlotskiSolve.dir/klotski.cpp.o


CMakeFiles/KlotskiSolve.dir/main.cpp.o: CMakeFiles/KlotskiSolve.dir/flags.make
CMakeFiles/KlotskiSolve.dir/main.cpp.o: ../main.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/zyp1/Klotski/algorithm/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/KlotskiSolve.dir/main.cpp.o"
	/usr/bin/c++   $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/KlotskiSolve.dir/main.cpp.o -c /home/zyp1/Klotski/algorithm/main.cpp

CMakeFiles/KlotskiSolve.dir/main.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/KlotskiSolve.dir/main.cpp.i"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/zyp1/Klotski/algorithm/main.cpp > CMakeFiles/KlotskiSolve.dir/main.cpp.i

CMakeFiles/KlotskiSolve.dir/main.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/KlotskiSolve.dir/main.cpp.s"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/zyp1/Klotski/algorithm/main.cpp -o CMakeFiles/KlotskiSolve.dir/main.cpp.s

CMakeFiles/KlotskiSolve.dir/main.cpp.o.requires:

.PHONY : CMakeFiles/KlotskiSolve.dir/main.cpp.o.requires

CMakeFiles/KlotskiSolve.dir/main.cpp.o.provides: CMakeFiles/KlotskiSolve.dir/main.cpp.o.requires
	$(MAKE) -f CMakeFiles/KlotskiSolve.dir/build.make CMakeFiles/KlotskiSolve.dir/main.cpp.o.provides.build
.PHONY : CMakeFiles/KlotskiSolve.dir/main.cpp.o.provides

CMakeFiles/KlotskiSolve.dir/main.cpp.o.provides.build: CMakeFiles/KlotskiSolve.dir/main.cpp.o


# Object files for target KlotskiSolve
KlotskiSolve_OBJECTS = \
"CMakeFiles/KlotskiSolve.dir/klotski.cpp.o" \
"CMakeFiles/KlotskiSolve.dir/main.cpp.o"

# External object files for target KlotskiSolve
KlotskiSolve_EXTERNAL_OBJECTS =

KlotskiSolve: CMakeFiles/KlotskiSolve.dir/klotski.cpp.o
KlotskiSolve: CMakeFiles/KlotskiSolve.dir/main.cpp.o
KlotskiSolve: CMakeFiles/KlotskiSolve.dir/build.make
KlotskiSolve: libTaskQueue.a
KlotskiSolve: libThreadPool.a
KlotskiSolve: CMakeFiles/KlotskiSolve.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/zyp1/Klotski/algorithm/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Linking CXX executable KlotskiSolve"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/KlotskiSolve.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/KlotskiSolve.dir/build: KlotskiSolve

.PHONY : CMakeFiles/KlotskiSolve.dir/build

CMakeFiles/KlotskiSolve.dir/requires: CMakeFiles/KlotskiSolve.dir/klotski.cpp.o.requires
CMakeFiles/KlotskiSolve.dir/requires: CMakeFiles/KlotskiSolve.dir/main.cpp.o.requires

.PHONY : CMakeFiles/KlotskiSolve.dir/requires

CMakeFiles/KlotskiSolve.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/KlotskiSolve.dir/cmake_clean.cmake
.PHONY : CMakeFiles/KlotskiSolve.dir/clean

CMakeFiles/KlotskiSolve.dir/depend:
	cd /home/zyp1/Klotski/algorithm/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/zyp1/Klotski/algorithm /home/zyp1/Klotski/algorithm /home/zyp1/Klotski/algorithm/build /home/zyp1/Klotski/algorithm/build /home/zyp1/Klotski/algorithm/build/CMakeFiles/KlotskiSolve.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/KlotskiSolve.dir/depend
