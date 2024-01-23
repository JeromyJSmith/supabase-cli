## supabase-login

Connect the Supabase CLI to your Supabase account by logging in with your [personal access token](https://supabase.com/dashboard/account/tokens).

Your access token is stored securely in [native credentials storage](https://github.com/zalando/go-keyring#dependencies). If native credentials storage is unavailable, it will be written to a plain text file at `~/.supabase/access-token`.

> In CI environments, you have the option to skip the login process. Instead of logging in, you can specify the `SUPABASE_ACCESS_TOKEN` environment variable in other commands to authenticate.

The Supabase CLI uses the stored token to access Management APIs for projects, functions, secrets, etc.
> **Note:** Skipping the login is particularly useful in CI workflows where you can set the `SUPABASE_ACCESS_TOKEN` as an environment variable for seamless command execution.
