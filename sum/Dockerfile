FROM pachyderm/job-shim
MAINTAINER erik.reppel@generalfusion.com
RUN sudo apt-get update

# Install python and pip
RUN sudo apt-get install -y \
    python-pip  \
    python-dev

# Install scipy tools
# RUN sudo apt-get install \
#     python-numpy         \
#     python-scipy         \
#     python-pandas        \
#     python-sympy         \
#     python-nose

# Install scikit-learn
# RUN pip install scikit-learn

# Install Tensorflow
# RUN sudo pip install --upgrade https://storage.googleapis.com/tensorflow/linux/cpu/tensorflow-0.8.0-cp27-none-linux_x86_64.whl

# Add local files
Add sum.py sum.py