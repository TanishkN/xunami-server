import os
import uuid
from pprint import pprint
from snaptrade_client import SnapTrade
from dotenv import load_dotenv
import os

load_dotenv()  # Load variables from .env file

# 1) Initialize a client with your clientID and consumerKey.
snaptrade = SnapTrade(
    consumer_key=os.environ["SNAPTRADE_CONSUMER_KEY"],
    client_id=os.environ["SNAPTRADE_CLIENT_ID"],
)

# 2) Check that the client is able to make a request to the API server.
api_response = snaptrade.api_status.check()
pprint(api_response.body)

# 3) Create a new user on SnapTrade
user_id = str(uuid.uuid4())
register_response = snaptrade.authentication.register_snap_trade_user(
    body={"userId": user_id}
)
pprint(register_response.body)
user_secret = register_response.body['userSecret']
# 4) Get a redirect URI. Users will need this to connect
# their brokerage to the SnapTrade server.
redirect_uri = snaptrade.authentication.login_snap_trade_user(
    query_params={"userId": user_id, "userSecret": user_secret}
)
print(redirect_uri.body)