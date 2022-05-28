import React, { useState } from "react";
import { Button, Text, TextInput, View } from "react-native";

type Props = {
    navigate(): void;
}

const Signup: React.FC<Props> = ({ navigate }) => {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    return (
        <View>
            <TextInput onChangeText={setUsername} value={username} />
            <TextInput secureTextEntry={true} onChangeText={setPassword} value={password} />
            <Button title="Create account"  onPress={navigate} />
        </View>
    );
}

export default Signup;