from datetime import datetime, timezone
import pandas as pd

def ms_to_datetime(ms):
  return datetime.utcfromtimestamp(float(ms) / 1000.)


def ms_to_pd_timestamp(ms):
  return pd.Timestamp(ms, unit='ms')


def timestamp_to_ms(timestamp):
  if isinstance(timestamp, datetime):
    return datetime_to_ms(timestamp)
  if not isinstance(timestamp, int):
    raise TypeError("couldn't parse {} as a timestamp".format(timestamp))
  return timestamp


def datetime_to_ms(dt):
  return int(dt.replace(tzinfo=timezone.utc).timestamp() * 1000)
