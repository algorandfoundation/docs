# compile TEAL code

from algosdk.v2client import algod
import os

# read file
def load_resource(res):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.join(dir_path, res)
    with open(path, "rb") as fin:
        data = fin.read()
    return data


try:
    # create an algod client
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_address = "http://localhost:4001"
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # int 1 - sample.teal
    # This code is meant for learning purposes only
    # It should not be used in production
    myprogram = "sample.teal"

    # read TEAL program
    data = load_resource(myprogram)
    source = data.decode("utf-8")

    # Compile TEAL program
    response = algod_client.compile(source)

    print("Response Result = ", response["result"])
    print("Response Hash = ", response["hash"])
except Exception as e:
    print(e)

# results should look similar to this:
# Response Result = ASABACI =
# Response Hash = KI4DJG2OOFJGUERJGSWCYGFZWDNEU2KWTU56VRJHITP62PLJ5VYMBFDBFE
