import React from "react";
import { Button, Text, View } from "react-native";
import List from "../models/list";

interface Props {
    list: List;
    setUpdate(): void;
}

const ListModal: React.FC<Props> = ({ list, setUpdate }) => {
    return (
        <View>
            <Text>
                {list.name}
            </Text>
            <Button title="Update list" onPress={ setUpdate } />
        </View>
    );
}

export default ListModal;