# compile teal code
from algosdk.v2client import algod

try:
    # create an algod client
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_address = "http://localhost:4001"
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # int 0 - sample.teal
    myprogram = "sample.teal"
    # read teal program
    data = open(myprogram, 'r').read()
    # compile teal program
    response = algod_client.compile(data)
    # print(response)
    print ("Response Result = ",response['result'])
    print("Response Hash = ",response['hash'])
except Exception as e:
    print(e)

# results should look similar to this:
# Response Result = ASABACI=
# Response Hash = KI4DJG2OOFJGUERJGSWCYGFZWDNEU2KWTU56VRJHITP62PLJ5VYMBFDBFE
