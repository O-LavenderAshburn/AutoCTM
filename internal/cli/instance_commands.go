package cli

import("fmt")

// AddLog adds a CT log to the instance’s configuration.
// The broker persists the log and notifies the instance to reload config.
func AddLog(instanceID string, url string) error {
    // validate input
    if url == "" {
        return fmt.Errorf("url cannot be empty")
    }
    if instanceID == "" {
        return fmt.Errorf("instanceID is required")
    }

    // call service layer
}


// RemoveLog removes a CT log from the instance’s configuration.
// The broker updates the database and notifies the instance to reload config.
func RemoveLog(instanceID string, url string) error {
    if instanceID == "" {
        return fmt.Errorf("instanceID is required")
    }
    if url == "" {
        return fmt.Errorf("url cannot be empty")
    }

}

//Pause the instance’s polling loop.
func Pause(instanceID string){

}

//Resume a paused instance.
func Resume(instanceID string){
	
}






