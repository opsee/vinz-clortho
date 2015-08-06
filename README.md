# Vinz Clortho

Gozer the Traveller, he will come in one of the pre-chosen forms. During the rectification of the Vuldronaii, the Traveller came as a large and moving Torb! Then, during the third reconciliation of the last of the Meketrex Supplicants they chose a new form for him... that of a Giant Sloar! Many Shubs and Zulls knew what it was to be roasted in the depths of the Sloar that day I can tell you.

## s3kms

s3kms is a simple utility for managing encrypted files in S3 for use with
AWS's Key Management Service (KMS). It supports two operations on an object
in S3: get and put. Assuming you have the appropriate key permissions in KMS:

`s3kms put -k alias/devkey -b mybucket -o someobject`

- Setup the appropriate encryption context in KMS
- Read from STDIN until EOF
- Encrypt the bytes read
- Put the encrypted data into the specified S3 bucket

`s3kms get -b mybucket -o someobject`

- Setup the appropriate encryption context in KMS
- Get the encrypted data from S3
- Decrypt the read data
- Write the encrypted data to STDOUT

s3kms is configured via environment variables or via the command line. See
`s3kms help` for more information.

* `AWS_KMS_KEY_ARN`: This is the ARN to the KMS key.
* `AWS_DEFAULT_REGION`: This is the region used for everything, basically.
* `AWS_ACCESS_KEY_ID`: The access key used should have the necessary privileges in KMS.
* `AWS_SECRET_ACCESS_KEY`: Private key for associated access key id.
* `AWS_ACCOUNT_ID`: Account ID for granting read access with s3kms put.

In practice, we use a different key for environments, as opposed to entity under
encryption. This seems to be the easiest way to configure things.
