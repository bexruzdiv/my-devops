FROM quay.io/ansible/awx-ee:latest
# Create a virtual environment and activate it

RUN curl -O https://raw.githubusercontent.com/kubernetes-sigs/kubespray/master/requirements.txt
RUN python3 -m venv kubespray-venv
RUN /bin/bash -c "source kubespray-venv/bin/activate"

# Copy your requirements.txt file and install dependencies
RUN pip3 install -U --user -r requirements.txt

# Set the entry point for the image
CMD ["/bin/bash"]



