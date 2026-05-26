package encryption

import (
	"crypto/cipher"

	"github.com/google/tink/go/kwp/subtle"
)

type AESKeyLength int

const (
	// DataKeyLength provides the length for data key in bytes.
	// It is recommended to use an authentication key with 32 or 64 bytes.
	// The data key length must be either
	// 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
	AES256KeyLength AESKeyLength = 32
	AES192KeyLength AESKeyLength = 24
	AES128KeyLength AESKeyLength = 16
)

var (
	masterKeyVarName = "LIFTBRIDGE_ENCRYPTION_KEY"
)

// LocalEncryptionHandler provides functionalities to load secret key
// from environment variables
type LocalEncryptionHandler struct {
	defaultDEK  []byte
	keyWrapper  *subtle.KWP
	blockCipher *cipher.AEAD
}

// NewLocalEncryptionHandler generates a new instance of LocalEncryptionHandler.
func NewLocalEncryptionHandler() (*LocalEncryptionHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Init key wrapper

// generateDEK generate a random AES data encryption key
func (handler *LocalEncryptionHandler) generateDEK() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (handler *LocalEncryptionHandler) wrapDEK(dek []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// use Tinker to wrap data key
	// https://github.com/google/tink/commit/22467ef7273d73b2d65e4b50310aab4af006bb7e
	return nil, nil
}

func (handler *LocalEncryptionHandler) unwrapDEK(wrappedDEK []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// use Tinker to unwrap data key
	// https://github.com/google/tink/commit/22467ef7273d73b2d65e4b50310aab4af006bb7e
	return nil, nil
}

func (handler *LocalEncryptionHandler) encryptData(dek []byte, plaintextData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// init cipher in GCM
	return nil, nil
}

// init nonce

func (handler *LocalEncryptionHandler) decryptData(dek []byte, encryptedData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get nonce

// decrypt the data

// Seal takes the message, performs the encryption and return
// the encrypted data along with the wrapped data
// The encoded message contains the first byte as size of the wrapped key,
// the wrapped key and finally the encrypted message.

// |  byte 0  |   byte 1   |   byte 2   |    ...   | byte (n +1 ) |    byte (n+2)  |  ... | byte (n + m + 2) |
// |----------|------------|------------|----------|--------------|----------------|------|------------------|
// | key size | key byte 0 | key byte 1 |      ... | key byte n   | message byte 0 |  ... |  message byte m  |

func (handler *LocalEncryptionHandler) Seal(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Generate a default Data Key (DEK) if not yet available
	return nil, nil
}

// encrypt the message

// wrap the data key

// concatenate:  key size | wrapped key | ciphertext

// allocate data sequence

// copy key size, wrapped key and cipher text to data sequence

// Read takes cipher text, performs the decryption and return
// the plaintext data
// The incoming byte array has the following structure:
// The first byte indicates the size of the wrapped key.
// The n next bytes contain the wrapped ky itself (n is the value of the first byte)
// The remaining bytes are the message itself.

// |  byte 0  |   byte 1   |   byte 2   |    ...   | byte (n +1 ) |    byte (n+2)  |  ... | byte (n + m + 2) |
// |----------|------------|------------|----------|--------------|----------------|------|------------------|
// | key size | key byte 0 | key byte 1 |      ... | key byte n   | message byte 0 |  ... |  message byte m  |

func (handler *LocalEncryptionHandler) Read(encryptedData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Decompose wrapped key and cypher text
	return nil, nil
}

// Decipher the message
