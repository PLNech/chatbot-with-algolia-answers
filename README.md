# Chatbot with Algolia Answers

This sample shows you how to leverage the Natural Language Processing features of [Algolia Answers](https://www.algolia.com/doc/guides/algolia-ai/answers/) by implement a chatbot connected to a knowledge base.

## Features
- 🧠 Natural Language Processing (NLP) using [Algolia Answers](https://www.algolia.com/doc/guides/algolia-ai/answers/).
- 🤖 Chatbot building made easy with [Dialogflow](https://cloud.google.com/dialogflow/).

## Demo (Try it yourself!)

[Access the demo](https://ni17w.sse.codesandbox.io/)


## How to run locally

This sample includes 2 server implementations in [Python](server/python) and [JavaScript (Node)](server/node).

The [client](client) is a single HTML page displaying the [Dialogflow Messenger](https://cloud.google.com/dialogflow/es/docs/integrations/dialogflow-messenger).

**1. Clone and configure the sample**

```
git clone https://github.com/algolia-samples/chatbot-with-algolia-answers
```

Copy the .env.example file into a file named .env in the folder of the server you want to use. For example:

```
cp .env.example server/go/.env
```

You will need an Algolia account in order to run the demo. If you don't have already an account, you can [create one for free](https://www.algolia.com/users/sign_up).

```bash
ALGOLIA_APP_ID=<replace-with-your-algolia-app-id>
```

**2. Create and populate your Algolia index**

Once your Algolia account and your Algolia application are setup, you will need to [create and populate an index](https://www.algolia.com/doc/guides/sending-and-managing-data/prepare-your-data/).

Alternatively, you populate your index with the same data used in the demo, [The World Health Organization COVID-19 FAQs](sample/who-covid-faq.json).

You can either upload your data directly from the [Algolia dashboard](https://www.algolia.com/doc/guides/sending-and-managing-data/send-and-update-your-data/how-to/importing-from-the-dashboard/) or by using one of our [API clients](https://www.algolia.com/developers/#integrations).

```bash
ALGOLIA_INDEX_NAME=<replace-with-your-algolia-index-name>
```

**3. Algolia Answers**

[Follow the instructions](https://www.algolia.com/doc/guides/algolia-ai/answers/#authentication) in order to setup Algolia Answers on your index.

```bash
ALGOLIA_API_KEY=<replace-with-your-algolia-api-key-with-algolia-answers-acl>
```

**4. Create and configure a Dialogflow agent**

[Dialogflow](https://cloud.google.com/dialogflow) is a Google Cloud solution for building virtual agents.
We will use it to quickly bootstrap our chatbot and use the [fulfillment](https://cloud.google.com/dialogflow/es/docs/fulfillment-overview) webhook feature to connect it to Algolia Answers.

1. Follow the Dialogflow guide to quickly create an empty agent.
2. Replace the default fallback intent with [this one](sample/dialogflow-default-fallback-intent.json).
3. Populate the `agent-id` variable in the [index file](client/index.html)

**5. Follow the server instructions on how to run**

Pick the server language you want and follow the instructions in the server folder README on how to run.

For example, if you want to run the Python server:

```
cd server/python # there's a README in this folder with instructions
python3 venv env
source env/bin/activate
pip3 install -r requirements.txt
export FLASK_APP=server.py
python3 -m flask run --port=4242
```

1. Expose your development server to the internet (by using [ngrok](https://ngrok.com/) or a similar solution).
2. Add the development server public URL to the Dialogflow webhook setting so it can use it for the [fulfillment](https://cloud.google.com/dialogflow/es/docs/fulfillment-overview) of the queries).

All set up!

## Authors
- [@cdenoix](https://twitter.com/cdenoix)
