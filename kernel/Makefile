#example make file for kernal
HIP_PATH = /opt/rocm/hip
HIPCC = $(HIP_PATH)/bin/hipcc 
TARGET_GPU = --targets gfx803 #gfx803 or gfx701 can be used too. If black 

all:
	# if no targets it will generate with all the different options
	# it will throw some warnings, but that is it
	$(HIPCC)  --genco example.cpp -o example.co
#If targeting a single machine use
#	$(HIPCC)  --genco $(TARGET_GPU) example.cpp -o example.co


