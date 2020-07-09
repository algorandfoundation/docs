import json
from algosdk import account, mnemonic

acct = account.generate_account()
address1 = acct[1]
print("Account 1")
print(address1)
mnemonic1 = mnemonic.from_private_key(acct[0])

print("Account 2")
acct = account.generate_account()
address2 = acct[1]
print(address2)
mnemonic2 = mnemonic.from_private_key(acct[0])

print("Account 3")
acct = account.generate_account()
address3 = acct[1]
print(address3)
mnemonic3 = mnemonic.from_private_key(acct[0])
print ("")
print("Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/")
print("copy off the following mnemonic code for use later")
print("")
print("mnemonic1 = \"{}\"".format(mnemonic1))
print("mnemonic2 = \"{}\"".format(mnemonic2))
print("mnemonic3 = \"{}\"".format(mnemonic3))


# terminal output should look similar to this
# Account 1
# KECA5Z2ZILJOH2ZG7OPKJ5KMFXP5XBAOC7H36TLPJOQI3KB5UIYUD5XTZU
# Account 2
# DWQ4IA7EK5BKHSPNCJBA5SOVU66TKGUCCGO2SHQCI5UB2JAO3G2GWXMGPA
# Account 3
# TABDMZ2EUTNOR3S74SJWW37DLHE7BDGS5XB5JPLFQ2VQVJOE2DXKX722VU

# Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at
# https: // bank.testnet.algorand.network/
# copy off the following mnemonic code for use later

# mnemonic1 = "consider round clerk soldier hurt dynamic floor video output spoon deliver virtual zoo inspire rubber doll nose warfare improve abstract recall choice size above actor"
# mnemonic2 = "boil explain enlist adapt science hub universe knife ghost scheme lazy payment must gas coconut forget goddess author filter civil tumble antique delay absorb lend"
# mnemonic3 = "place elbow thumb bid taste strong sting tube swarm comic wave dinosaur congress sword zebra need proud primary brief rotate story pilot garbage abstract black"
