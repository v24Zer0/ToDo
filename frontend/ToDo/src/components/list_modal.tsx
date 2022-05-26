import React from "react";
import { Alert, Button, Text, View } from "react-native";
import List from "../models/list";

interface Props {
    list: List;
    setUpdate(): void;
}

const ListModal: React.FC<Props> = ({ list, setUpdate }) => {
    return (
        <View>
            <Text>
                Update list {list.name}
            </Text>
            <Text>
                {list.name}
            </Text>
            <Button title="Update list" onPress={ setUpdate } />
            <Button title="Delete list" onPress={() => Alert.alert("List deleted")} />
        </View>
    );
}

export default ListModal;