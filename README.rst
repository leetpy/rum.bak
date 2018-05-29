===
Rum
===

Install
=======

.. code-block:: console

    $ go get -u github.com/leetpy/rum

Run
===

.. code-block:: console

    $ cd $GOPATH/src/github.com/leetpy/rum
    $ go run main.go

string_to_sign = 'GET\n/iaas/\naccess_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1'
h = hmac.new(secret_access_key, digestmod=sha256)
h.update(string_to_sign)
sign = base64.b64encode(h.digest()).strip()
signature = urllib.quote_plus(sign)
