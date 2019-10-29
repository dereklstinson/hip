HIP_PATH = /opt/rocm/hip
CFLAGS = -I$(HIP_PATH)/include -D__HIP_PLATFORM_HCC__ -D__HIP_VDI__ 
LDFLAGS = -L$(HIP_PATH)/lib -lhip_hcc


all:

	gcc -O3  main.c $(CFLAGS) $(LDFLAGS) -o main.out