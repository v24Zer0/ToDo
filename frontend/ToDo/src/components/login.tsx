import React, { useState } from "react";
import { Alert, Button, Text, TextInput, View } from "react-native";

type Props = {
    navigate(): void;
}

const Login: React.FC<Props> = ({ navigate }) => {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    return (
        <View>
            <TextInput onChangeText={setUsername} placeholder="Username" value={username} />
            <TextInput onChangeText={setPassword} secureTextEntry={true} placeholder="Password" value={password} />
            <Button title="login" onPress={navigate} />
        </View>
    );
}

export default Login;