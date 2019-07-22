import os
import sys
from os.path import dirname, join

from dotenv import load_dotenv

dotenv_path = join(dirname(__file__), '.env') # Explicitly only load the .env file in the same directory as this file
load_dotenv(dotenv_path)

# List all required environment variables here.
# If they exist, they will be loaded as 'settings.MY_ENV'.
# If they are missing, program will tell you what's missing and exit
REQUIRED_ENVS = [
    "WEB3_PROVIDER_URL", "BENEATH_BASE_URL", "BENEATH_PROJECT",
    "BENEATH_PROJECT_KEY", "BENEATH_PROJECT_STREAM"
]

# List all optional environment variables here.
# They will be loaded if they exist, but no error will result if they're missing
OPTIONAL_ENVS = ["SOME_OPTIONAL_VARIABLE"]

"""
Check that a ENV exists and load it
"""
def require_env(env_key):
    if os.environ.get(env_key):
        return os.environ.get(env_key)
    else:
        print(f"ERROR - Missing environment variable: {env_key}")
        sys.exit()

# Load all required ENVs
for env in REQUIRED_ENVS:
    globals()[env] = require_env(env)

# Load all optional ENVs
for env in OPTIONAL_ENVS:
    globals()[env] = os.environ.get(env)
