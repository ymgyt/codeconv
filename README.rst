========
codeconv
========

codeconv is utility command to convert(encode/decode) common data format.

Usage
=====

.. code-block:: bash

   # base64
   $ codeconv base64 encode "hello gopher" | codeconv base64 decode -
   hello gopher