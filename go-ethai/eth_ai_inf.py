import os
import sys
import tensorflow as tf
import numpy as np
import re

os.environ['TF_CPP_MIN_LOG_LEVEL']='2'

def bitfield(n):
        return [int(digit) for digit in bin(n)[2:]]

def toBinary(n):
        return ''.join(str(1 & int(n) >> i) for i in range(64)[::-1])


#difficulty_t = np.genfromtxt("/home/naksir/WORK/go-ethai/test_input_data.txt",  delimiter="," )
#data = difficulty_t[0]
data1 = toBinary(int(sys.argv[1]))
#print data1

data = [] 
for i in range(0,64):
        data.append( int(data1[i]))

x = tf.placeholder(tf.float32, [64]) #any number of row is good
w = tf.Variable(tf.random_uniform([64], -1.0, 1.0))
b = tf.Variable(tf.random_uniform([64], -1.0, 1.0))
y = tf.nn.sigmoid( x*w + b)

saver = tf.train.Saver()
init_op = tf.global_variables_initializer()

with tf.Session() as sess:
    sess.run(init_op)
    saver.restore(sess, "/home/naksir/WORK/go-ethai/sigmoid.bin")
    prediction = sess.run(tf.round(y), feed_dict={x:data})
    
    #print prediction
    result = 0
    for i in range(0,64):
        result = result | (int(prediction[i]) << (63-i))
    print result
