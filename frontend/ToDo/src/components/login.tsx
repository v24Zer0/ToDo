import React, { useState } from "react";
import { Alert, Button, Text, TextInput, View } from "react-native";

const Login = () => {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    return (
        <View>
            <TextInput onChangeText={setUsername} placeholder="Username" value={username} />
            <TextInput onChangeText={setPassword} placeholder="Password" value={password} />
            <Button title="login" onPress={() => Alert.alert("login button pressed")} />
        </View>
    );
}

export default Login;