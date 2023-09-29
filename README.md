# Paperspace Provider for DevPod

This is a [Paperspace](https://paperspace.com) provider for [DevPod](https://github.com/dwarvesf/devpod), initially forked from https://github.com/dirien/devpod-provider-scaleway. This repository uses Paperspace's [machines API](https://docs.paperspace.com/core/api-reference/machines) to provision the machines.

![](https://i.imgur.com/TdpCAz5.png)

## Environment Variables

There are 2 environment variables that you will need to be aware of:

1. `PPS_API_KEY`: The API key you generate for your account on Paperspace. You will be prompted for this when you add the provider.

    To acquire your API key, you can get it in your workspace settings on Paperspace. For more info, you can look at their docs [here](https://docs.paperspace.com/account-management/account/security/api-keys).

    ![](https://docs.paperspace.com/assets/images/security-api-key-2-1ee96c963c8e029f4594c02eeb40bacc.png)

2. `SSH_FOLDER`: The directory of your Paperspace SSH keys. **By default, it is your `~/.ssh` folder.**

    DevPod will generate your keys and save them to your `SSH_FOLDER`. The SSH client library uses `github.com/loft-sh/devpod/pkg/ssh`, which will try to find __`id_devpod_rsa` and `id_devpod_rsa.pub`__ in your `SSH_FOLDER`. If it can't find those files, it will generate its own.

    As of v0.0.15, this provider will automatically authorize the SSH keys in your `SSH_FOLDER` to your machine.

## Getting started

You can install and use this provider by running the following commands:

```sh
devpod provider add github.com/dwarvesf/devpod-provider-paperspace
devpod provider use github.com/dwarvesf/devpod-provider-paperspace
```

### Trying it out

You can use Microsoft's [`vscode-remote-try-node`](https://github.com/microsoft/vscode-remote-try-node)'s example app to try out the provider. To keep it interactive, we'll use the `openvscode` IDE to run VSCode in the browser. You can run the following command to try it out:

```sh
devpod up github.com/microsoft/vscode-remote-try-node --ide openvscode
```

It should take a bit while it is provisioning the machine for you, but you will be emailed by Paperspace that you have created a new machine. You can reference an example output from the command here: https://app.warp.dev/block/pgQghAkqstTdyCevilplVa

You should then see a new window open to your new localhost server, running VSCode in the browser.

![](https://i.imgur.com/H3EXRub.png)

## Customizing the workspace

You can customize the environment variables after installation by updating the DevPod provider options like so:
```sh
devpod provider use paperspace -o PPS_MACHINE_TYPE="GPU+"
```

With this, there are a few environment variables you should be aware of so that you can set and customize your machine type and OS:

1. `PPS_MACHINE_TYPE`: The type of machine you want to use.

    The default machine type that gets provisioned uses `C4`, which offers 2 CPUs and 4GB of RAM. If you would like to choose a different machine type, you can reference the tables below:

    | Label | CPUs | RAM (GB) | GPU  |
    | ----- | ---- | -------- | ---- |
    | C4    | 2    | 4        | None |
    | C5    | 4    | 8        | None |
    | C6    | 8    | 16       | None |
    | C7    | 12   | 30       | None |
    | C8    | 16   | 60       | None |
    | C9    | 24   | 120      | None |
    | C10   | 32   | 244      | None |

    You also have the option to use machines with GPUs:

    | Label     | CPUs | RAM (GB) | GPU            |
    | --------- | ---- | -------- | -------------- |
    | GPU+      | 8    | 30       | Quadro M4000   |
    | P4000     | 8    | 30       | Quadro P4000   |
    | P5000     | 8    | 30       | Quadro P5000   |
    | P6000     | 8    | 30       | Quadro P6000   |
    | RTX4000   | 8    | 30       | Quadro RTX4000 |
    | RTX5000   | 8    | 30       | Quadro RTX5000 |
    | P5000x2   | 16   | 60       | Quadro P5000   |
    | P6000x2   | 16   | 60       | Quadro P6000   |
    | RTX5000x2 | 16   | 60       | Quadro RTX5000 |
    | P4000x2   | 16   | 60       | Quadro P4000   |

    You can find more info about the available machine types [here](https://docs.paperspace.com/core/compute/machine-types).

    To edit your machine type for your workspace, manually set the environment variable when using the provider:

    ```sh
    devpod provider use paperspace -o PPS_MACHINE_TYPE="GPU+"
    ```

2. `PPS_MACHINE_TEMPLATE`: The machine template and OS you want to use. The default template that is used when DevPod first installs the provider is `t0nspur5`, which runs a headless Ubuntu 22.04 server with a default disk size of 50GB.

    Paperspace comes with a handful of machine templates. They don't document it, unfortunately, so you can reference the template IDs from the table below.

    | ID       | Agent Type     | Operating System Label              | Default Size (GB) |
    | -------- | -------------- | ----------------------------------- | ----------------- |
    | tnr2oh1m | WindowsDesktop | Windows 10 (Server 2022) - Licensed | 50                |
    | t9taj00e | LinuxHeadless  | CentOS 7 Server                     | 50                |
    | t04azgph | LinuxHeadless  | Ubuntu 18.04 Server                 | 50                |
    | t0nspur5 | LinuxHeadless  | Ubuntu 22.04 Server                 | 50                |
    | tkni3aa4 | LinuxHeadless  | Ubuntu 20.04 Server                 | 50                |
    | tpi7gqht | LinuxHeadless  | Ubuntu 22.04 K8s Worker             | 50                |


    You can set the `PPS_MACHINE_TEMPLATE` with the ID of the operating system you wish to use, like so:

    ```sh
    devpod provider use paperspace -o PPS_MACHINE_TEMPLATE="t9taj00e"
    ```

3. `PPS_DISK_SIZE`: The disk size of the virtual machine you want to use. The default disk size is 50GB.

    You can set the provider option just like with the other environment variables. The value you set will be represented in GBs:
    ```sh
    devpod provider use paperspace -o PPS_DISK_SIZE="100"
    ```
