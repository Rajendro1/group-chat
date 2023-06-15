// Build a Realtime group chat app in Golang using WebSockets


const domElement = document.querySelector(".chat__app-container");

class App extends React.Component {
    constructor() {
        super();
        this.state = {
            messages: [],
        }
        this.webSocketConnection = null;
    }

    componentDidMount() {
        this.setWebSocketConnection();
        this.subscribeToSocketMessage();
    }

    setWebSocketConnection() {
        const username = prompt("What's Your name");
        //console.log("the username: ",username)
        if (window["WebSocket"]) {
            const socketConnection = new WebSocket("ws://" + document.location.host + "/ws/" + username);
            console.log("username is: ",username)
            this.webSocketConnection = socketConnection;
        }
    }


    subscribeToSocketMessage = () => {
        if (this.webSocketConnection === null) {
            return;
        }

        this.webSocketConnection.onclose = (evt) => {
            const messages = this.state.messages;
            messages.push({
                message: 'Your Connection is closed.',
                type: 'announcement'
            })
            this.setState({
                messages
            });
        };

        this.webSocketConnection.onmessage = (event) => {
            try {
                const socketPayload = JSON.parse(event.data);
                switch (socketPayload.eventName) {
                    case 'join':
                        if (!socketPayload.eventPayload) {
                            return
                        }

                        this.setState({
                            messages: [
                                ...this.state.messages,
                                ...[{
                                    message: `${socketPayload.eventPayload} joined the chat`,
                                    type: 'announcement'
                                }]
                            ]
                        });

                        break;
                    case 'disconnect':
                        if (!socketPayload.eventPayload) {
                            return
                        }
                        this.setState({
                            messages: [
                                ...this.state.messages,
                                ...[{
                                    message: `${socketPayload.eventPayload} left the chat`,
                                    type: 'announcement'
                                }]
                            ]
                        });
                        break;

                    case 'message response':

                        if (!socketPayload.eventPayload) {
                            return
                        }

                        const messageContent = socketPayload.eventPayload;
                        const sentBy = messageContent.username ? messageContent.username : 'An unnamed fellow'
                        const actualMessage = messageContent.message;

                        const messages = this.state.messages;
                        messages.push({
                            message: actualMessage,
                            username: `${sentBy} says:`,
                            type: 'message'
                        })

                        this.setState({
                            messages
                        });

                        break;

                    default:
                        break;
                }
            } catch (error) {
                console.log(error)
                console.warn('Something went wrong while decoding the Message Payload')
            }
        };
    }

    handleKeyPress = (event) => {
        try {
            if (event.key === 'Enter') {
                if (!this.webSocketConnection) {
                    return false;
                }
                if (!event.target.value) {
                    return false;
                }

                this.webSocketConnection.send(JSON.stringify({
                    EventName: 'message',
                    EventPayload: event.target.value
                }));

                event.target.value = '';
            }
        } catch (error) {
            console.log(error)
            console.warn('Something went wrong while decoding the Message Payload')
        }
    }

    getChatMessages() {
        return (
            <div class="message-container">
                {
                    this.state.messages.map(m => {
                        return (
                            <div class="message-payload">
                                {m.username && <span class="username">{m.username}</span>}
                                <span class={`message ${m.type === 'announcement' ? 'announcement' : ''}`}>{m.message}</span>
                            </div>
                        )
                    })
                }
            </div>
        );
    }

  processValueAndSendToGo(value) {
        // Process the value using Babel-transpiled code
        const processedValue = processValue(value);
      
        // Send the processed value to the Go server
        fetch('/sendProcessedValue', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ processedValue }),
        })
          .then(response => response.json())
          .then(data => {
            console.log('Processed value sent to Go successfully:', data);
            // Handle any additional logic or UI updates
          })
          .catch(error => {
            console.error('Error sending processed value to Go:', error);
            // Handle error case
          });
      }
      
      // Babel-transpiled code for processing the value
 processValue(value) {
        // Your Babel-transpiled code logic here
        // Example: Convert value to uppercase
        return value.toUpperCase();
      }
      
      // Call the function with the desired value
    //   const  = 'Hello, Go!';
    //   processValueAndSendToGo(value);
      
      
   
    render() {
        return (
            <>
                {this.getChatMessages()}
                <input type="text" id="message-text" size="64" autofocus placeholder="Type Your message" onKeyPress={this.handleKeyPress} />
            </>
        );
    }
}

ReactDOM.render(<App />, domElement)
