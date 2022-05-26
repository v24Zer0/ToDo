import React from "react";
import { Button } from "react-native";

type Props = {
    navigate(): void; 
};

const UserButton: React.FC<Props> = ({ navigate }) => {
    return (
        <Button title="User" onPress={navigate} />
    );
}

export default UserButton;