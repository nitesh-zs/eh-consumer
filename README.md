# EventHub-Consumer

EventHub consumer is a go application intended to be used by QA team for testing of the event-collector project. As of now, this project specifically deals with NEW IMF events only. 

## How to use the project:

- Clone this repository.
- Create a file `.local.env` under the `configs` folder.
- Copy the internally shared eventhub credentials in `.local.env`.
- The binaries present in `bin` folder can be used to run the consumer application.
- Run the binary based on your system's architecture. `amd64` for intel chip macs and `arm64` for m1 macs.

### Command to consume IMF-Update events:
``` bin/eventhub_consumer_arm64 load -config=imf_config.json ```

The events data will be saved in `imf_out.json` file.

### Command to consume BOH-Check events:
``` bin/eventhub_consumer_arm64 load -config=boh_config.json ```

The events data will be saved in `boh_out.json` file.