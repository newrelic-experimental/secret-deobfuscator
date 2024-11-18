<a href="https://opensource.newrelic.com/oss-category/#new-relic-experimental"><picture><source media="(prefers-color-scheme: dark)" srcset="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/dark/Experimental.png"><source media="(prefers-color-scheme: light)" srcset="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/Experimental.png"><img alt="New Relic Open Source experimental project banner." src="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/Experimental.png"></picture></a>

# secret-deobfuscator

 A utility for reversing the obfuscation of sensitive information performed using the [New Relic CLI](https://docs.newrelic.com/docs/infrastructure/host-integrations/installation/secrets-management/#newrelic-cli-obfuscation)

## Installation
Pre-built binaries are available for Linux, MacOS, and Windows through the [releases page](https://github.com/newrelic-experimental/secret-deobfuscator/releases). Alternatively you can build the project from source, which requires a local installation of [Go](https://go.dev/doc/install) :
```sh
$ wget https://raw.githubusercontent.com/newrelic-experimental/secret-deobfuscator/refs/heads/main/nrDeobfuscate.go
$ go build ./nrDeobfuscate.go
```

## Usage
 1. Download [the release](https://github.com/newrelic-experimental/secret-deobfuscator/releases) that matches your environment, or build the binary from source. Linux and MacOS can utilize `wget` for this process
 ```
 $ wget https://github.com/newrelic-experimental/secret-deobfuscator/releases/download/v1.0.0/nrDeobfuscate-[Linux|Darwin]
 ```
 2. Open a shell (Bash, Terminal, PowerShell, etc.) and `cd` into the directory where your binary is kept
 3. Use `chmod +x ./nrDeobfuscate-[release]` to make it executable (unless you're using Windows)
 4. Pass the obfuscated string (your '_secret_') and the key you originally used to obfuscate your payload to the binary via arguments:
 ```sh
 # Linux
 $ ./nrDeobfuscate-Linux -key '<value>' -secret '<value>'

# MacOS
$ ./nrDeobfuscate-Darwin -key '<value>' -secret '<value>'

# Windows
$ ./nrDeobfuscate-Windows.exe -key '<value>' -secret '<value>'
 ```

## Example
Using the [New Relic CLI](https://docs.newrelic.com/docs/infrastructure/host-integrations/installation/secrets-management/#newrelic-cli-obfuscation) we can obfuscate our payload:
```sh
$ newrelic agent config obfuscate --value '{"Username":"username@example.com","Password":"Sup3r$ecr3tP@ssw0rd!"}' --key 'SecretKey'

{
  "obfuscatedValue": "KEc2AQAGJQQUNkdZUBAHLhcXMggGMgAMKggJPwBNEQoZaUlbAwQQARIbOQFbaUcwBxVHOUEcMBdQBjU0OBYOYxcHU0cJ"
}
```

Now we can use the deobfuscator to return the obfuscated value back to readable JSON:
```sh
$ ./nrDeobfuscate-Darwin -secret 'KEc2AQAGJQQUNkdZUBAHLhcXMggGMgAMKggJPwBNEQoZaUlbAwQQARIbOQFbaUcwBxVHOUEcMBdQBjU0OBYOYxcHU0cJ' -key 'SecretKey'

'{"Username":"username@example.com","Password":"Sup3r$ecr3tP@ssw0rd!"}'
```

## Planned changes:
N/A

## Support

Requests for support can be filed as an [Issue](https://github.com/newrelic-experimental/secret-deobfuscator/issues).

## Contributing
We encourage your contributions to improve `secret-deobfuscator`! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project.
If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company,  please drop us an email at opensource@newrelic.com.

**A note about vulnerabilities**

As noted in our [security policy](../../security/policy), New Relic is committed to the privacy and security of our customers and their data. We believe that providing coordinated disclosure by security researchers and engaging with the security community are important means to achieve our security goals.

If you believe you have found a security vulnerability in this project or any of New Relic's products or websites, we welcome and greatly appreciate you reporting it to New Relic through [HackerOne](https://hackerone.com/newrelic).

## License
`secret-deobfuscator` is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.
