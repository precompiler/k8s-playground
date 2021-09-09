package com.precompiler;


import com.fasterxml.jackson.databind.JsonNode;
import io.fabric8.kubernetes.api.model.Status;
import io.fabric8.kubernetes.api.model.admission.v1.AdmissionReview;
import io.fabric8.kubernetes.api.model.admission.v1.AdmissionReviewBuilder;
import lombok.extern.slf4j.Slf4j;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.ArrayList;
import java.util.List;

/**
 * @author Richard Li
 */
@Slf4j
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
@Path("/")
public class WebhookResource {
    @Path("/validate")
    @POST
    public Response validate(JsonNode payload) {
        List<String> errors = new ArrayList<>();
        log.info("validating request: {}", payload);
        JsonNode uidNode = payload.at("/request/uid");
        if (uidNode.isMissingNode()) {
            errors.add("Missing request UID");
        }
        JsonNode replicasNode = payload.at("/request/object/spec/replicas");
        if (replicasNode.isMissingNode()) {
            errors.add("Cannot find replicas info");
        }
        int replicas = replicasNode.asInt();
        if (replicas > 1) {
            errors.add("Deployment replicas cannot be greater than ONE");
        }
        AdmissionReview response = null;
        if (errors.size() > 0) {
            Status status = new Status();
            status.setCode(400);
            status.setMessage(errors.toString());
            response = new AdmissionReviewBuilder().withNewResponse().withAllowed(Boolean.FALSE).withUid(uidNode.asText()).withStatus(status).endResponse().build();
        } else {
            response = new AdmissionReviewBuilder().withNewResponse().withAllowed(Boolean.TRUE).withUid(uidNode.asText()).endResponse().build();
        }

        return Response.ok(response).build();
    }
}
