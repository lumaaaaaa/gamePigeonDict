# gamePigeonDict
All iOS GamePigeon dictionaries dumped from the app's .ipa.

## Layout
- `dicts`: Alternate language dictionaries. They're stored seperately, since they're unencrypted 
in the .ipa.
- `encrypted`: Encrypted English language dictionaries.
    - `gp_en.dat`: Core dictionary. Contains words that are at most 6 characters long.
    - `gp_en_2.dat`: Extended dictionary. Contains all words in `gp_en.dat` and more.
    - `gp_en_c.dat`: Censored dictionary. Contains words that will be replaced with asterisks when played.
- `decrypted`: Plaintext result of decrypting the encrypted dictionaries with this repository's Go script. Censored 
  directory is omitted, to prevent indexing in my GitHub.
- `main.go`: Go script to decrypt the dictionaries.
- `go.mod`: Go module file.

## Decryption
The dictionaries are encrypted with AES-256-CBC. The key is `T6wfOZgP0QgasdsgT6wfOZgP0Qgasdsg` and the IV is 
`T6wfOZgP0Qgasdsg`. Decryption is simple with the provided Go script. 

## Usage
1. Clone the repository.
2. Run `go run main.go` in the repository's root directory.
3. The decrypted dictionaries will be outputted to the `decrypted` directory.