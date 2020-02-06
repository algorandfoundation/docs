title: Algorand Smart Contracts

Algorand Smart Contracts are small programs written in an assembly-like language that can be used as a replacement for signatures within a transaction. This language is named Transaction Execution Approval Langauge or TEAL. TEAL programs have one primary function and that is to return true or false and are used to analyze and approve transactions. Algorand transactions are signed with either a private key or a multisig set of private keys. Using ASC1, transactions can be signed with a TEAL program. This is called a logic signature. ASC provides two modes of operation for TEAL logic to operate as a LogicSignature, that are discussed in Usage Modes<LINK>.




To start using Alogrand Smart Contracts it is important to understand several aspects. Primarily this can be separated into learning TEAL and how to call TEAL from either the command line `goal` tool or by calling it from any of the SDKs.

To learn more about the TEAL language specification see the TEAL Specification Reference documentation<LINK>. To get a simplified understanding of how TEAL is processed see the TEAL overview guide<LINK>.

To compile and use TEAL programs using `goal` see the the goal TEAL walkthrough<LINK>. 

To start working with TEAL with any of the SDKs, see the ASC1 SDK guide<LINK>.




