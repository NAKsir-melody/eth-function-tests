import os
import tensorflow as tf
import numpy as np

os.environ['TF_CPP_MIN_LOG_LEVEL']='2'

input_size = sum(1 for line in open('/home/naksir/WORK/go-ethai/training_input_data.txt'))
print input_size
input_size = input_size -1

difficulty = np.genfromtxt("/home/naksir/WORK/go-ethai/training_input_data.txt",  delimiter=",", max_rows=input_size )
#print difficulty
nonce = np.genfromtxt("/home/naksir/WORK/go-ethai/training_label_data.txt",  delimiter="," , max_rows=input_size )
#print nonce
difficulty_t = np.genfromtxt("/home/naksir/WORK/go-ethai/test_input_data.txt",  delimiter="," )
#print difficulty_t
nonce_t = np.genfromtxt("/home/naksir/WORK/go-ethai/test_label_data.txt",  delimiter="," )
#print nonce_t


# input data
x = tf.placeholder(tf.float32, [64]) #any number of row is good
w = tf.Variable(tf.random_uniform([64], -1.0, 1.0))
b = tf.Variable(tf.random_uniform([64], -1.0, 1.0))
y = tf.nn.sigmoid( x*w + b)
# output
y_target= tf.placeholder(tf.float32, [64])

#set training option
#error function = RMSE
#optimize method = minimize cost
cost = tf.reduce_mean(tf.square(y - y_target))
train_step = tf.train.GradientDescentOptimizer(0.1).minimize(cost)

init = tf.global_variables_initializer()
sess = tf.Session()
sess.run(init)

#training
for i in range(input_size):
    xs = difficulty[i]
    ys = nonce[i]
    sess.run(train_step, feed_dict={x: xs , y_target: ys})   

inf = tf.equal(tf.round(y), y_target)
accuracy = tf.reduce_mean(tf.cast(inf, tf.float32))

#print single results
result = [None] * 10
for j in range(10):
    xt = difficulty_t[j]
    yt = nonce_t[j]
    inf_result = sess.run(inf, feed_dict={x: xt, y_target: yt} )
    result[j] = sess.run(accuracy, feed_dict={x: xt, y_target: yt} )
    print yt
    print "----"
    print inf_result
    print "===="

avg = 0;
for k in range(10):
    avg += float(result[k])/10
print avg

#save model
saver = tf.train.Saver()
save_path = saver.save(sess,"./sigmoid.bin")






    


