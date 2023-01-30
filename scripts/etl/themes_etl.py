import json 
import pandas as pd


SETS_FILEPATH = '.../static/json/lego_sets.json'
THEMES_FILEPATH = '.../static/json/lego_themes.json'


def extract(fp: str) -> dict:
	with open(fp, encoding="utf-8") as f:
		return json.load(f)


def transform(data: list) -> pd.DataFrame:
	df = pd.DataFrame(data).reset_index() 
	df.rename(columns={'index': 'id'}, inplace=True)
	df['id'] = df['id'].apply(lambda x: x + 1)

	return df


def load(df: pd.DataFrame, output_fp: str):
	df.to_json(output_fp, indent=4)


if __name__ == '__main__':
	load(
		transform(
			extract(THEMES_FILEPATH)
		),
		output_fp='../backend/static/json/lego_themes_modified.json'
	)